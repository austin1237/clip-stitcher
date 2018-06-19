package uploader

import (
	"fmt"
	"log"
	"net/http"
	"time"

	youtube "google.golang.org/api/youtube/v3"
)

func uploadToYouTube(video Video, client *http.Client) error {
	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating YouTube client: %v", err)
	}
	vidTitle := fmt.Sprintf("%v clip highlights %v", video.ChannelName, time.Now().Format("01-02-2006"))
	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       vidTitle,
			Description: video.VideoDescription,
			CategoryId:  "22",
		},
		Status: &youtube.VideoStatus{PrivacyStatus: "unlisted"},
	}

	call := service.Videos.Insert("snippet,status", upload)

	_, err = call.Media(video.FileStream).Do()
	if err != nil {
		return err
	}
	return nil
}
