package twitch

import (
	"errors"
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

	tclips = filterOutOldClips(tclips)
	tclips = filterOutNonVod(tclips)
	tclips, err = filterOutOverlap(tclips)

	if err != nil {
		return preparedClips, err
	}

	if len(tclips.Clips) < 2 {
		err = errors.New("Not enough clips left to combine after filter")
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

func filterOutNonVod(clips twitchAPIResp) twitchAPIResp {
	clipsWithVod := twitchAPIResp{}
	for _, clip := range clips.Clips {
		if clip.Vod.URL != "" && clip.Vod.ID != "" {
			clipsWithVod.Clips = append(clipsWithVod.Clips, clip)
		}
	}
	return clipsWithVod
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
