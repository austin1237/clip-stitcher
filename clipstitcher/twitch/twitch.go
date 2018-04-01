package twitch

import (
	"fmt"
	"time"
)

type TwitchService struct {
	streamName   string
	clientID     string
	desiredCount int
}

func NewTwitchService(name string, count int, id string) *TwitchService {
	service := new(TwitchService)
	service.desiredCount = count
	service.clientID = id
	service.streamName = name
	return service
}

func (tService TwitchService) GetClips() (PreparedClips, error) {
	preparedClips := PreparedClips{}
	tclips, err := getClips(tService.streamName, tService.desiredCount, tService.clientID)
	if err != nil {
		return preparedClips, err
	}
	tclips, err = filterOutOverlap(tclips)

	if err != nil {
		return preparedClips, err
	}

	pageURLs := make([]string, 0)

	for _, clip := range tclips.Clips {
		pageURLs = append(pageURLs, clip.URL)
	}

	clipHTMLS, err := getClipHTML(pageURLs)
	if err != nil {
		return preparedClips, err
	}
	clipSrcs, err := getSrcsFromHTML(clipHTMLS)
	if err != nil {
		return preparedClips, err
	}
	preparedClips.VideoLinks = clipSrcs
	preparedClips.VideoDescription = generateDescription(tclips)
	return preparedClips, err
}

func generateDescription(clips twitchAPIResp) string {
	description := ""
	duration := 0.00
	startTime := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
	endTime := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
	for _, clip := range clips.Clips {
		line := fmt.Sprintf("%v %v \n", clip.Title, clip.URL)
		description = description + line
		duration += clip.Duration
		endTime = endTime.Add(time.Second * time.Duration(clip.Duration))
	}

	duration = endTime.Sub(startTime).Minutes()
	description = description + fmt.Sprintf("Total lengh of the video should be %v long", duration)
	return description
}

func getClipHTML(clipUrls []string) ([]string, error) {
	clipHTMLs := []string{}
	desiredCount := len(clipUrls)
	htmlResponses := make(chan asyncString, desiredCount)

	for _, url := range clipUrls {
		go asyncGetClipHTML(url, htmlResponses)
	}
	// Wait for the Responses from the scraper
	fmt.Println("Waiting for Scraper to finish")
	for response := range htmlResponses {
		err := response.err
		if err != nil {
			close(htmlResponses)
			return []string{}, err
		}
		clipHTMLs = append(clipHTMLs, response.value)
		if len(clipHTMLs) == desiredCount {
			close(htmlResponses)
		}
	}
	return clipHTMLs, nil
}

func getSrcsFromHTML(clipHTMLs []string) ([]string, error) {
	clipSrcs := []string{}
	for _, clipHTML := range clipHTMLs {
		clipSrc, err := findVidSrcInHTML(clipHTML)
		if err != nil {
			return []string{}, err
		}
		clipSrcs = append(clipSrcs, clipSrc)
	}
	return clipSrcs, nil
}
