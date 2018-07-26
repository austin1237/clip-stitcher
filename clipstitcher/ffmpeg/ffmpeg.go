package ffmpeg

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Service struct {
	Cmd        *exec.Cmd
	FileStream io.ReadCloser
	Logs       io.ReadCloser
}

// Here's a blog where I found the ffmpeg command to combine videos with a transcode
//http://www.bugcodemaster.com/article/concatenate-videos-using-ffmpeg
func buildFFmpegcommand(clipLinks []string) string {
	cmdString := "ffmpeg"
	inputs := ""
	streamOptions := ""
	for index, link := range clipLinks {
		iStr := strconv.Itoa(index)
		inputs = inputs + "-i " + link + " "
		streamOptions = streamOptions + "[" + iStr + ":v:0] [" + iStr + ":a:0] "
	}
	cmdString = "ffmpeg " + inputs + "-filter_complex \"" + streamOptions + "concat=n=" + strconv.Itoa(len(clipLinks)) + ":v=1:a=1 [v] [a]\" "
	cmdString = cmdString + "-map [v] -map [a] -q:v 0 -q:a 0 -r 60 -f avi -"
	return cmdString
}

func NewFFmpegService(clipLinks []string) (*Service, error) {
	service := new(Service)
	cmdTxt := buildFFmpegcommand(clipLinks)
	service.Cmd = exec.Command("bash", "-c", cmdTxt)
	fileStream, err := service.Cmd.StdoutPipe()
	if err != nil {
		return service, err
	}
	service.FileStream = fileStream
	logs, err := service.Cmd.StderrPipe()
	if err != nil {
		return service, err
	}
	service.Logs = logs
	return service, nil
}

func (service Service) Start() (*bytes.Buffer, error) {
	fmt.Println("ffmpeg start hit")
	buffer := &bytes.Buffer{}
	bufferChan := make(chan *bytes.Buffer)
	errsChan := make(chan error)
	defer service.Logs.Close()
	defer service.FileStream.Close()
	go checkOutputErrs(service.Logs, errsChan)
	go bufferFileStream(service.FileStream, errsChan, bufferChan)
	err := service.Cmd.Start()
	if err != nil {
		return buffer, err
	}
	for i := 0; i < 2; i++ {
		err := <-errsChan
		if err != nil {
			service.Cmd.Process.Kill()
			return buffer, err
		}
	}
	buffer = <-bufferChan
	return buffer, nil
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
			done <- errors.New("Error found in output: " + output)
		}
	}
	done <- nil
}

func bufferFileStream(fileStream io.Reader, errChan chan error, bufferChan chan *bytes.Buffer) {
	buffer := bytes.Buffer{}
	_, err := buffer.ReadFrom(fileStream)
	if err != nil && err != io.EOF {
		fmt.Println("err is " + err.Error())
		buffer.Reset()
		errChan <- err
	} else {
		errChan <- nil
		bufferChan <- &buffer
	}
}
