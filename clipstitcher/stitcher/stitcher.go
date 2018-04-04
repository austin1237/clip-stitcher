package stitcher

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
)

var Logs []byte

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

func StitchClips(clipLinks []string) (io.ReadCloser, error) {
	cmd := buildFFmpegcommand(clipLinks)
	log.Println(cmd)
	ffmpeg := exec.Command("bash", "-c", cmd)

	fileStream, err := ffmpeg.StdoutPipe()
	if err != nil {
		return nil, err
	}
	logs, err := ffmpeg.StderrPipe()
	if err != nil {
		return nil, err
	}
	go keepsLogsInMemory(logs)
	err = ffmpeg.Start()
	if err != nil {
		fmt.Println("error starting " + cmd)
		return nil, err
	}

	return fileStream, nil

}

func keepsLogsInMemory(stdErr io.ReadCloser) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stdErr)
	Logs = buf.Bytes()
}
