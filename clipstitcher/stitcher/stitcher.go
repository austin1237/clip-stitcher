package stitcher

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/ffmpeg"
	"github.com/user/clipstitcher/uploader"
)

func StitchAndUpload(clipMessage consumer.ClipMessage, ytAuth string) error {
	retryCount := 3
	buffer := &bytes.Buffer{}
	ffmpegService, err := ffmpeg.NewFFmpegService(clipMessage.VideoLinks)
	if err != nil {
		return err
	}
	for retry := 1; retry <= retryCount; retry++ {
		buffer, err = ffmpegService.Start()
		if err == nil {
			break
		} else if retry == retryCount {
			return errors.New("Max retries hit with ffmpeg")
		}
		fmt.Println("Retrying ffmpeg")
	}
	video := uploader.Video{
		FileStream:       buffer,
		VideoDescription: clipMessage.VideoDescription,
		ChannelName:      clipMessage.ChannelName,
	}
	err = uploader.Upload(video, ytAuth)
	return err
}
