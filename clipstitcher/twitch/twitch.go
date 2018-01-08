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
	clipSrcs, err := getClipSrcs(pageUrls)
	return clipSrcs, err
}

func getClipSrcs(clipUrls []string) ([]string, error) {
	clipSrcs := []string{}
	desiredCount := len(clipUrls)
	scrapResponses := make(chan asyncString, desiredCount)

	for _, url := range clipUrls {
		go scrapVidSrcOnPage(url, scrapResponses)
	}
	// Wait for the Responses from the scraper
	fmt.Println("Waiting for Scraper to finish")
	for response := range scrapResponses {
		err := response.err
		if err != nil {
			close(scrapResponses)
			return []string{}, err
		}
		clipSrcs = append(clipSrcs, response.value)
		if len(clipSrcs) == desiredCount {
			close(scrapResponses)
		}
	}
	fmt.Println("Scraper finished")
	return clipSrcs, nil

}
