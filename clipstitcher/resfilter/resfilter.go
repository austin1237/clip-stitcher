package resfilter

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

var ErrNoHighResoultionClip = errors.New("No clips found with high enough resoultion")

type VideoResGroup struct {
	High        []string
	SuperMedium []string
	Medium      []string
	Low         []string
}

func getVideoRes(clipLink string) (string, error) {
	cmdTxt := "ffprobe -v error -select_streams v:0 -show_entries stream=width,height -of csv=s=x:p=0 " + clipLink
	cmd := exec.Command("bash", "-c", cmdTxt)
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	resString := string(bytes)
	resString = strings.TrimSpace(resString)
	return resString, nil
}

func groupByRes(videoLinks []string) (VideoResGroup, error) {
	resGroup := VideoResGroup{}
	for _, link := range videoLinks {
		res, err := getVideoRes(link)
		if err != nil {
			fmt.Println(err.Error())
			return VideoResGroup{}, err
		}
		fmt.Println(res)

		if res == "1920x1080" {
			resGroup.High = append(resGroup.High, link)
		} else if res == "1600x900" {
			resGroup.SuperMedium = append(resGroup.SuperMedium, link)
		} else if res == "1280x720" {
			resGroup.Medium = append(resGroup.Medium, link)
		} else {
			resGroup.Low = append(resGroup.Low, link)
		}
	}
	return resGroup, nil
}

func FilterOutLowRes(videoLinks []string) ([]string, error) {
	groupedClips, err := groupByRes(videoLinks)
	if err != nil {
		return []string{}, err
	}
	bestClips, err := getBestResoultion(groupedClips)
	if err != nil {
		return []string{}, err
	}
	return bestClips, nil
}

func getBestResoultion(groupedClips VideoResGroup) ([]string, error) {
	if len(groupedClips.High) > 1 {
		return groupedClips.High, nil
	} else if len(groupedClips.SuperMedium) > 1 {
		return groupedClips.SuperMedium, nil
	} else if len(groupedClips.Medium) > 1 {
		return groupedClips.Medium, nil
	}
	err := ErrNoHighResoultionClip
	return []string{}, err
}
