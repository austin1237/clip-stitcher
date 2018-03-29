package uploader

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	youtube "google.golang.org/api/youtube/v3"
)

func uploadToYouTube(fileStream io.ReadCloser, client *http.Client, videoDescirption string, channelName string) {
	flag.Parse()
	defer fileStream.Close()

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}

	vidTitle := fmt.Sprintf("%v clip highlights %v", channelName, time.Now().Format("01-02-2006"))
	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       vidTitle,
			Description: videoDescirption,
			CategoryId:  "22",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
	}

	call := service.Videos.Insert("snippet,status", upload)

	response, err := call.Media(fileStream).Do()
	if err != nil {
		log.Fatalf("Error making YouTube API call: %v", err)
	}
	fmt.Printf("Upload successful! Video ID: %v\n", response.Id)
}
