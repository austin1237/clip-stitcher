package main

import (
	"fmt"
	"os"
	"time"

	"github.com/user/clipstitcher/consumer"
	"github.com/user/clipstitcher/resfilter"
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
		fmt.Println("exiting due to error")
		fmt.Printf("Error: %+v", err)
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
	start := time.Now()
	consumerService, err := consumer.NewConsumerService(consumerEndpoint, consumerURL)
	logAndExit(err)
	clipMessage, err := consumerService.GetMessage()
	fmt.Println("Message found for " + clipMessage.ChannelName)
	fmt.Println(len(clipMessage.VideoLinks))
	logAndExit(err)
	filteredVideoLinks, err := resfilter.FilterOutLowRes(clipMessage.VideoLinks)
	logAndExit(err)
	fmt.Println("Resoultions filter for " + clipMessage.ChannelName)
	clipMessage.VideoLinks = filteredVideoLinks
	err = stitcher.StitchAndUpload(clipMessage, youtubeAuth)
	logAndExit(err)
	fmt.Println("Video stitching finished for " + clipMessage.ChannelName)
	err = consumerService.DeleteMessage(clipMessage)
	logAndExit(err)
	elapsed := time.Since(start)
	fmt.Println("message deleted for " + clipMessage.ChannelName)
	fmt.Printf("total execution took %s for %s \n", elapsed, clipMessage.ChannelName)
	os.Exit(0)
}
