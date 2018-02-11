package stitcher

import (
	"fmt"
	"io"
	"os/exec"
	"strconv"
)

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
	cmdString = cmdString + "-map [v] -map [a] -frag_duration 3600 -f mp4 -"
	return cmdString
}

func StitchClips(clipLinks []string) (io.ReadCloser, error) {

	cmd := buildFFmpegcommand(clipLinks)
	ffmpeg := exec.Command("bash", "-c", cmd)

	fileStream, err := ffmpeg.StdoutPipe()
	if err != nil {
		return nil, err
	}
	err = ffmpeg.Start()
	if err != nil {
		fmt.Println("error starting " + cmd)
		return nil, err
	}

	return fileStream, nil

}
