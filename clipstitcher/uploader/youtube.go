package uploader

import (
	"fmt"
	"net/http"
	"time"

	"github.com/pkg/errors"
	youtube "google.golang.org/api/youtube/v3"
)

func uploadToYouTube(video Video, client *http.Client) error {
	service, err := youtube.New(client)
	if err != nil {
		err = errors.New(err.Error())
		return err
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
		err = errors.New(err.Error())
		return err
	}
	return nil
}
