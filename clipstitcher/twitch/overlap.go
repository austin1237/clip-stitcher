package twitch

import (
	"errors"
	"strconv"
	"strings"
)

func filterOutOverlap(clips twitchClips) twitchClips {
	return clips
}

// func loopsThroughClips(clips twitchClips) {
// 	clipMap := make(map[int][]clipTime)
// 	for _, clip := range clips.Clips {
// 		clipMap[clip.Vod.ID] = clip

// 	}
// }

func getTimeFromURL(url string) (clipTime, error) {
	var cTime clipTime
	timeStr := strings.Split(url, "=")[1]
	hourStr := strings.Split(timeStr, "h")[0]
	timeStr = strings.Split(timeStr, "h")[1]
	minuteStr := strings.Split(timeStr, "m")[0]
	timeStr = strings.Split(timeStr, "m")[1]
	secondsStr := strings.Split(timeStr, "s")[0]
	hours, err := strconv.Atoi(hourStr)
	if err != nil {
		return cTime, errors.New("invalid hours found in vod url")
	}
	minutes, err := strconv.Atoi(minuteStr)
	if err != nil {
		return cTime, errors.New("invalid minutes found in vod url")
	}
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return cTime, errors.New("invalid seconds found in vod url")
	}
	cTime.Seconds = seconds
	cTime.Minutes = minutes
	cTime.Hours = hours
	return cTime, nil
}
