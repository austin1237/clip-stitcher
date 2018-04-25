package main

import (
	"fmt"
	"os"
	"strings"
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
	youtubeAuth string
)

func logAndExit(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("exiting due to error")
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

	youtubeAuth = os.Getenv("YOUTUBE_AUTH")
	if twitchChannelName == "" {
		fmt.Println("YOUTUBE_AUTH ENV var was not set.")
		os.Exit(1)
	}
}

func main() {
	fmt.Println("clip sticher started")
	start := time.Now()
	twitchService := twitch.NewTwitchService(twitchChannelName, 10, twitchClientID)
	preparedClips, err := twitchService.GetClips()
	logAndExit(err)
	fmt.Println("starting ffmpeg")
	ffmpegReader, err := stitcher.StitchClips(preparedClips.VideoLinks)
	logAndExit(err)
	fmt.Println("starting upload")
	err = uploader.Upload(ffmpegReader, youtubeAuth, preparedClips.VideoDescription, twitchChannelName)
	if err != nil {
		logAndExit(err)
	}
	fmt.Println("upload finished")
	elapsed := time.Since(start)
	sitcherOutput := strings.Replace(string(stitcher.Logs), "%", "%%", -1)
	fmt.Println(sitcherOutput)
	fmt.Println("total execution time took", elapsed)
	os.Exit(0)
}
