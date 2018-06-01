package stitcher

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/ffmpeg"
	"github.com/user/clipstitcher/uploader"
)

func StichAndUpload(clipMessage consumer.ClipMessage, ytAuth string) error {
	videoLinks := clipMessage.VideoLinks
	dupFileRd, dupFileWr := io.Pipe()
	ffmpegService, err := ffmpeg.NewFFmpegService(videoLinks)
	if err != nil {
		return err
	}
	teeFileStream := io.TeeReader(ffmpegService.FileStream, dupFileWr)
	video := uploader.Video{
		FileStream:       teeFileStream,
		VideoDescription: clipMessage.VideoDescription,
		ChannelName:      clipMessage.ChannelName,
	}

	errsChan := make(chan error)
	defer close(errsChan)
	go checkOutputErrs(ffmpegService.Logs, errsChan)
	go makeSureDataIsStreaming(dupFileRd, errsChan)
	go uploader.Upload(video, dupFileWr, ytAuth, errsChan)
	ffmpegService.Cmd.Start()
	for i := 0; i < 3; i++ {
		err := <-errsChan
		if err != nil {
			return err
		}
	}
	return nil
}

func makeSureDataIsStreaming(stream io.Reader, mainDone chan error) {
	p := make([]byte, 4)
	noDataCounter := 0
	errsChan := make(chan error)
	defer close(errsChan)

	// If data is recevied lower the counter
	go func(counter *int, readerDone chan error) {
		for {
			_, err := stream.Read(p)
			if err != nil {
				if err == io.EOF {
					break
				}
				readerDone <- err
			}
			currentCount := *counter
			if currentCount > 0 {
				*counter = currentCount - 1
			}
		}
		readerDone <- nil
	}(&noDataCounter, errsChan)

	// Up the counter every X seconds throw err if counter reaches 3
	go func(counter *int, tickerDone chan error) {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			currentCount := *counter + 1
			*counter = currentCount
			if currentCount > 3 {
				err := errors.New("No data has been streamed in awhile")
				tickerDone <- err
			}
		}
	}(&noDataCounter, errsChan)

	// Returns either nil or err
	err := <-errsChan
	mainDone <- err
}

func checkOutputErrs(stream io.ReadCloser, done chan error) {
	defer stream.Close()
	output := ""
	r := bufio.NewReader(stream)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			done <- err
		}
		lineStr := string(line)
		lineStrLower := strings.ToLower(lineStr)
		output = output + lineStr + "\n"
		if strings.Contains(lineStrLower, "error") {
			fmt.Println("Error found in output " + lineStr)
		}
	}
	done <- nil
}
