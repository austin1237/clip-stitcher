package twitch

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func filterOutOverlap(clips twitchAPIResp) (twitchAPIResp, error) {
	var finishedClips twitchAPIResp
	overlapMap := make(map[int][]overlap)
	for _, clip := range clips.Clips {
		cTime, err := getTimeFromURL(clip.Vod.URL)
		if err != nil {
			return finishedClips, err
		}
		cDuration := getDurationFromTime(cTime, clip.Duration)
		tOverlap := overlap{
			VodID:     clip.Vod.ID,
			StartTime: cDuration.StartTime,
			EndTime:   cDuration.EndTime,
		}

		if !isClipOverlaping(overlapMap, tOverlap) && len(finishedClips.Clips) <= 10 {
			overlapMap[clip.Vod.ID] = append(overlapMap[clip.Vod.ID], tOverlap)
			finishedClips.Clips = append(finishedClips.Clips, clip)
		}
	}
	return finishedClips, nil
}

// Determine if clip is overlaping here
func isClipOverlaping(clipMap map[int][]overlap, tOverlap overlap) bool {
	return false
}

func getDurationFromTime(cTime clipTime, duration float64) clipDuration {
	var cDuration clipDuration
	startTime := time.Date(1970, time.January, 1, 8, 0, 0, 0, time.UTC)
	startTime = startTime.Add(time.Hour*time.Duration(cTime.Hours) + time.Minute*time.Duration(cTime.Minutes) + time.Second*time.Duration(cTime.Seconds))
	endTime := startTime.Add(time.Second * time.Duration(duration))
	cDuration.StartTime = startTime
	cDuration.EndTime = endTime
	return cDuration
}

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
