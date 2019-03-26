package twitch

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func filterOutOldClips(clips twitchAPIResp) twitchAPIResp {
	var filteredClips twitchAPIResp
	before := len(clips.Clips)
	for _, clip := range clips.Clips {
		now := time.Now()
		timeDiff := now.Sub(clip.CreatedAt)
		hourDiff := timeDiff.Hours()
		fmt.Println(hourDiff)
		// If a clip was created longer than 1 day ago it will be ignored.
		if hourDiff < 24 {
			filteredClips.Clips = append(filteredClips.Clips, clip)
		}
	}
	after := len(filteredClips.Clips)
	fmt.Println(before)
	fmt.Println(after)
	return filteredClips
}

func filterOutOverlap(clips twitchAPIResp) (twitchAPIResp, error) {
	var finishedClips twitchAPIResp
	overlapMap := make(map[string][]twitchDuration)
	for _, clip := range clips.Clips {
		cTime, err := getTimeFromURL(clip.Vod.URL)
		if err != nil {
			return finishedClips, err
		}
		cDuration := getDurationFromTime(cTime, clip.Duration)
		tOverlap := twitchDuration{
			VodID:     clip.Vod.ID,
			StartTime: cDuration.StartTime,
			EndTime:   cDuration.EndTime,
		}

		if !isClipOverlaping(overlapMap[clip.Vod.ID], tOverlap) && len(finishedClips.Clips) <= 10 {
			overlapMap[clip.Vod.ID] = append(overlapMap[clip.Vod.ID], tOverlap)
			finishedClips.Clips = append(finishedClips.Clips, clip)
		}
	}
	return finishedClips, nil
}

// Determine if clip is overlaping here
func isClipOverlaping(checkedDurations []twitchDuration, newDuration twitchDuration) bool {
	for _, checkedDuration := range checkedDurations {
		newStart := newDuration.StartTime
		newEnd := newDuration.EndTime
		oldStart := checkedDuration.StartTime
		oldEnd := checkedDuration.EndTime
		startDifference := oldStart.Sub(newEnd).Seconds()
		endDifference := oldEnd.Sub(newStart).Seconds()

		// https://stackoverflow.com/questions/325933/determine-whether-two-date-ranges-overlap?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
		// (StartA <= EndB) && (EndA >= StartB)
		if startDifference <= 0 && endDifference >= 0 {
			return true
		}
	}
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

	if len(strings.Split(timeStr, "h")) == 2 {
		hourStr := strings.Split(timeStr, "h")[0]
		timeStr = strings.Split(timeStr, "h")[1]
		hours, err := strconv.Atoi(hourStr)
		if err != nil {
			return cTime, errors.New("invalid hours found in vod url")
		}
		cTime.Hours = hours
	}

	if len(strings.Split(timeStr, "m")) == 2 {
		minuteStr := strings.Split(timeStr, "m")[0]
		timeStr = strings.Split(timeStr, "m")[1]
		minutes, err := strconv.Atoi(minuteStr)
		if err != nil {
			return cTime, errors.New("invalid minutes found in vod url")
		}
		cTime.Minutes = minutes
	}

	if len(strings.Split(timeStr, "s")) == 2 {
		secondsStr := strings.Split(timeStr, "s")[0]
		seconds, err := strconv.Atoi(secondsStr)
		if err != nil {
			return cTime, errors.New("invalid seconds found in vod url")
		}
		cTime.Seconds = seconds
	}

	return cTime, nil
}
