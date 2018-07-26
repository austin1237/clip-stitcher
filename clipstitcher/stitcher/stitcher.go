package stitcher

import (
	"bytes"
	"fmt"
	"time"

	"github.com/pkg/errors"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/ffmpeg"
	"github.com/user/clipstitcher/uploader"
)

func StitchAndUpload(clipMessage consumer.ClipMessage, ytAuth string) error {
	retryCount := 3
	buffer := &bytes.Buffer{}
	transStart := time.Now()
	for retry := 1; retry <= retryCount; retry++ {
		ffmpegService, err := ffmpeg.NewFFmpegService(clipMessage.VideoLinks)
		if err != nil {
			return errors.New(err.Error())
		}
		buffer, err = ffmpegService.Start()
		if err == nil {
			break
		} else if retry == retryCount {
			return errors.New("Max retries hit with ffmpeg")
		}
		fmt.Printf("Error: %+v", err)
		fmt.Println("Retrying ffmpeg")
		buffer = &bytes.Buffer{}
	}
	transTotal := time.Since(transStart)
	fmt.Printf("ffmpeg finished, took %s for %s \n", transTotal, clipMessage.ChannelName)
	uploadStart := time.Now()
	video := uploader.Video{
		FileStream:       buffer,
		VideoDescription: clipMessage.VideoDescription,
		ChannelName:      clipMessage.ChannelName,
	}
	err := uploader.Upload(video, ytAuth)
	if err != nil {
		return errors.New(err.Error())
	}
	uploadTotal := time.Since(uploadStart)
	fmt.Printf("upload finished, took %s for %s \n", uploadTotal, clipMessage.ChannelName)
	return nil
}
