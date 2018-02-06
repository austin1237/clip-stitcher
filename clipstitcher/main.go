package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/user/clipstitcher/stitcher"
	"github.com/user/clipstitcher/twitch"
	"github.com/user/clipstitcher/uploader"
)

var (
	// TWITCH ENV VARIABLES
	twitchClientID    string
	twitchChannelName string

	// YOUTUBE ENV VARIABLES
	youtubeClientID     string
	youtubeSecret       string
	youtubeAccessToken  string
	youtubeRefreshToken string
	youtubeExpiry       time.Time
)

func logAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	twitchClientID = os.Getenv("TWITCH_CLIENT_ID")
	if twitchClientID == "" {
		fmt.Println("TWITCH_CLIENT_ID ENV var was not set.")
		os.Exit(1)
	}

	twitchChannelName = os.Getenv("TWITCH_CHANNEL_NAME")
	if twitchChannelName == "" {
		fmt.Println("TWITCH_CHANNEL_NAME ENV var was not set.")
		os.Exit(1)
	}

	youtubeClientID = os.Getenv("YOUTUBE_CLIENT_ID")
	if youtubeClientID == "" {
		fmt.Println("YOUTUBE_CLIENT_ID ENV var was not set.")
		os.Exit(1)
	}

	youtubeSecret = os.Getenv("YOUTUBE_SECRET")
	if youtubeSecret == "" {
		fmt.Println("YOUTUBE_SECRET ENV var was not set.")
		os.Exit(1)
	}

	youtubeAccessToken = os.Getenv("YOUTUBE_ACCESS_TOKEN")
	if youtubeAccessToken == "" {
		fmt.Println("YOUTUBE_ACCESS_TOKEN ENV var was not set.")
		os.Exit(1)
	}

	youtubeRefreshToken = os.Getenv("YOUTUBE_REFRESH_TOKEN")
	if youtubeRefreshToken == "" {
		fmt.Println("YOUTUBE_REFRESH_TOKEN ENV var was not set.")
		os.Exit(1)
	}

	youtubeExpiryStr := os.Getenv("YOUTUBE_EXPIRY")
	if youtubeExpiryStr == "" {
		fmt.Println("YOUTUBE_EXPIRY ENV var was not set.")
		os.Exit(1)
	}
	date, err := time.Parse(time.RFC3339, youtubeExpiryStr)
	if err != nil {
		fmt.Println("YOUTUBE_EXPIRY is not a valid timestamp")
		os.Exit(1)
	}
	youtubeExpiry = date

}

func main() {
	fmt.Println("clip sticher started")
	start := time.Now()
	twitchService := twitch.NewTwitchService("drdisrespectlive", 10, twitchClientID)
	videoLinks, err := twitchService.GetVideoLinks()
	logAndExit(err)
	fmt.Println("videoLinks in main", videoLinks)
	fmt.Println("starting ffmpeg")
	ffmpegReader, err := stitcher.StitchClips(videoLinks)
	logAndExit(err)
	uploader.Upload(ffmpegReader, youtubeClientID, youtubeSecret, youtubeAccessToken, youtubeRefreshToken, youtubeExpiry)
	elapsed := time.Since(start)
	fmt.Println("exiting")
	log.Printf("total execution time took %s", elapsed)
	os.Exit(0)
}
