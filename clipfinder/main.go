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

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/user/clipfinder/producer"
	"github.com/user/clipfinder/twitch"
)

var (
	//ENV VARIABLES
	twitchClientID    string
	twitchChannelName string
	producerName      string
	producerEndpoint  string
)

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

	producerName = os.Getenv("PRODUCER_NAME")
	if producerName == "" {
		fmt.Println("PRODUCER_NAME ENV var was not set.")
		os.Exit(1)
	}

	producerEndpoint = os.Getenv("PRODUCER_ENDPOINT")
}

//Test
func HandleRequest(ctx context.Context, event events.S3Event) (string, error) {
	twitchService := twitch.NewTwitchService(twitchChannelName, 10, twitchClientID)
	preparedClips, err := twitchService.GetClips()
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	producerService, err := producer.NewProducerService(producerEndpoint, producerName)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	err = producerService.SendMessage(preparedClips.VideoLinks, preparedClips.VideoDescription, twitchChannelName)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}
	return "Message Sent", nil
}

func main() {
	lambda.Start(HandleRequest)
}
