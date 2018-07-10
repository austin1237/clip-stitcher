package main

import (
	"fmt"
	"os"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/stitcher"
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
	fmt.Println("clip stitcher started")
	consumerService, err := consumer.NewConsumerService(consumerEndpoint, consumerURL)
	logAndExit(err)
	clipMessage, err := consumerService.GetMessage()
	logAndExit(err)
	fmt.Println("Message found for " + clipMessage.ChannelName)
	err = stitcher.StitchAndUpload(clipMessage, youtubeAuth)
	logAndExit(err)
	fmt.Println("Video stitching finished for " + clipMessage.ChannelName)
	err = consumerService.DeleteMessage(clipMessage)
	logAndExit(err)
	fmt.Println("message deleted for " + clipMessage.ChannelName)
	os.Exit(0)
}
