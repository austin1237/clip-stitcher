package stitcher

import (
	"fmt"
	"os/exec"
	"strconv"
)

// Here's a blog where I found the ffmpeg command to combine videos with a transcode
//http://www.bugcodemaster.com/article/concatenate-videos-using-ffmpeg
func buildFFmpegcommand(output string, numOfClips int) string {
	cmdString := "ffmpeg"
	inputs := ""
	streamOptions := ""
	for i := 0; i < numOfClips; i++ {
		iStr := strconv.Itoa(i)
		inputs = inputs + "-i ./tmpClips/clip" + iStr + ".mp4 "
		streamOptions = streamOptions + "[" + iStr + ":v:0] [" + iStr + ":a:0] "
	}
	cmdString = "ffmpeg " + inputs + "-filter_complex \"" + streamOptions + "concat=n=" + strconv.Itoa(numOfClips) + ":v=1:a=1 [v] [a]\" "
	cmdString = cmdString + "-map [v] -map [a] -y " + output
	return cmdString
}

func StitchClips(outputLocation string, numOfClips int) {

	cmd := buildFFmpegcommand(outputLocation, numOfClips)
	fmt.Println(cmd)
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			fmt.Println(string(exiterr.Stderr))
		}
	}
	fmt.Printf("output from ffmpeg is %s\n", out)

}
