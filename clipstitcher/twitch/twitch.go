package twitch

import (
	"fmt"
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

func (tService TwitchService) GetVideoLinks() ([]string, error) {
	pageUrls, err := getClipPageUrls(tService.streamName, tService.desiredCount, tService.clientID)
	if err != nil {
		return []string{}, err
	}
	clipHTMLS, err := getClipHTML(pageUrls)
	if err != nil {
		return []string{}, err
	}
	clipSrcs, err := getSrcsFromHTML(clipHTMLS)
	return clipSrcs, err
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
