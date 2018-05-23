package main

import (
	"fmt"
	"os"
	"time"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/stitcher"
	"github.com/user/clipstitcher/uploader"
)

var (
	// YOUTUBE ENV VARIABLES
	youtubeAuth      string
	consumerEndpoint string
	consumerURL      string
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

	consumerURL = os.Getenv("CONSUMER_URL")
	if consumerURL == "" {
		fmt.Println("CONSUMER_URL ENV var was not set.")
		os.Exit(1)
	}

	consumerEndpoint = os.Getenv("CONSUMER_ENDPOINT")
}

func main() {
	fmt.Println("clip sticher started")
	start := time.Now()
	consumerService, err := consumer.NewConsumerService(consumerEndpoint, consumerURL)
	logAndExit(err)
	clipMessage, err := consumerService.GetMessage()
	fmt.Println("message data recieved")
	logAndExit(err)
	fmt.Println("starting ffmpeg")
	ffmpegReader, err := stitcher.StitchClips(clipMessage.VideoLinks)
	logAndExit(err)
	fmt.Println("starting upload")
	err = uploader.Upload(ffmpegReader, youtubeAuth, clipMessage.VideoDescription, clipMessage.ChannelName)
	if err != nil {
		fmt.Println(string(stitcher.Logs))
		logAndExit(err)
	}
	fmt.Println("upload finished")
	elapsed := time.Since(start)
	err = consumerService.DeleteMessage(clipMessage)
	fmt.Println("message deleted")

	logAndExit(err)
	fmt.Println("total execution time took", elapsed)
	os.Exit(0)
}
