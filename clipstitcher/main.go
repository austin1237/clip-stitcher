package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/stitcher"
	"github.com/user/clipstitcher/uploader"
)

var (
	// YOUTUBE ENV VARIABLES
	youtubeAuth      string
	consumerEndpoint string
	consumerName     string
)

func logAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting due to error")
		os.Exit(1)
	}
}

func init() {
	youtubeAuth = os.Getenv("YOUTUBE_AUTH")
	if youtubeAuth == "" {
		fmt.Println("YOUTUBE_AUTH ENV var was not set.")
		os.Exit(1)
	}

	consumerName = os.Getenv("CONSUMER_NAME")
	if consumerName == "" {
		fmt.Println("CONSUMER_NAME ENV var was not set.")
		os.Exit(1)
	}

	consumerEndpoint = os.Getenv("CONSUMER_ENDPOINT")
}

func main() {
	fmt.Println("clip sticher started")
	start := time.Now()
	consumerService, err := consumer.NewConsumerService(consumerEndpoint, consumerName)
	logAndExit(err)
	clipMessage, err := consumerService.GetMessage()
	logAndExit(err)
	fmt.Println("starting ffmpeg")
	ffmpegReader, err := stitcher.StitchClips(clipMessage.VideoLinks)
	logAndExit(err)
	fmt.Println("starting upload")
	err = uploader.Upload(ffmpegReader, youtubeAuth, clipMessage.VideoDescription, clipMessage.ChannelName)
	if err != nil {
		logAndExit(err)
	}
	fmt.Println("upload finished")
	elapsed := time.Since(start)
	sitcherOutput := strings.Replace(string(stitcher.Logs), "%", "%%", -1)
	fmt.Println(sitcherOutput)
	err = consumerService.DeleteMessage(clipMessage)
	logAndExit(err)
	fmt.Println("total execution time took", elapsed)
	os.Exit(0)
}
