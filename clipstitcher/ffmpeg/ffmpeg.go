package ffmpeg

import (
	"io"
	"os/exec"
	"strconv"
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
