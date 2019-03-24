// Compile with:
// docker run --rm -v "$PWD":/go/src/handler lambci/lambda:build-go1.x sh -c 'dep ensure && go build handler.go'

// Run with:
// docker run --rm -v "$PWD":/var/task lambci/lambda:go1.x handler '{"Records": []}'

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/user/clipfinder/producer"
	"github.com/user/clipfinder/twitch"
)

var (
	//ENV VARIABLES
	twitchClientID     string
	twitchChannelNames string
	producerURL        string
	producerEndpoint   string
)

func init() {
	twitchClientID = os.Getenv("TWITCH_CLIENT_ID")
	if twitchClientID == "" {
		fmt.Println("TWITCH_CLIENT_ID ENV var was not set.")
		os.Exit(1)
	}

	twitchChannelNames = os.Getenv("TWITCH_CHANNEL_NAME")
	if twitchChannelNames == "" {
		fmt.Println("TWITCH_CHANNEL_NAME ENV var was not set.")
		os.Exit(1)
	}

	producerURL = os.Getenv("PRODUCER_URL")
	if producerURL == "" {
		fmt.Println("PRODUCER_URL ENV var was not set.")
		os.Exit(1)
	}

	producerEndpoint = os.Getenv("PRODUCER_ENDPOINT")
}

func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {
	channels := strings.Split(twitchChannelNames, ",")
	for _, channelName := range channels {
		twitchService := twitch.NewTwitchService(channelName, 10, twitchClientID)
		preparedClips, err := twitchService.GetClips()
		if err != nil {
			log.Fatal(err.Error())
			return "", err
		}
		pService := producer.NewProducerService(producerEndpoint, producerURL)
		err = pService.SendMessage(preparedClips.VideoSlugs, preparedClips.VideoDescription, channelName)
		if err != nil {
			log.Fatal(err.Error())
			return "", err
		}
		fmt.Println("Message Sent for " + channelName)
	}
	return "All channel messages sent", nil
}

func main() {
	lambda.Start(HandleRequest)
}
