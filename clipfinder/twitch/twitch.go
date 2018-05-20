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

	clipSlugs := make([]string, 0)

	for _, clip := range tclips.Clips {
		clipSlugs = append(clipSlugs, clip.Slug)
	}

	preparedClips.VideoSlugs = clipSlugs
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
