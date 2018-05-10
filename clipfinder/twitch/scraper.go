package twitch

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func asyncGetClipHTML(clipURL string, response chan asyncString) {
	resp, err := http.Get(clipURL)
	if err != nil {
		response <- asyncString{"", err}
		return
	}
	defer resp.Body.Close()
	if err != nil {
		response <- asyncString{"", err}
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	response <- asyncString{string(body), nil}
}

func findVidSrcInHTML(clipHTML string) (string, error) {

	qualityOptionsTxt, err := findQualtyOptions(clipHTML)

	if err != nil {
		return "", err
	}

	clips, err := convertStringToClips(qualityOptionsTxt)
	if err != nil {
		return "", err
	}
	clipSrc, err := pickBestQuailty(clips)
	if err != nil {
		return "", err
	}
	return clipSrc, nil
}

func findQualtyOptions(HTMLText string) (string, error) {
	qualityOptions := ""
	// Loops through the multiline js script
	for _, line := range strings.Split(strings.TrimSuffix(HTMLText, "\n"), "\n") {
		if strings.Contains(line, "quality_options") {
			qualityOptions = strings.Replace(line, " ", "", -1)
			break
		}
	}
	if qualityOptions == "" {
		err := errors.New("quality_options not found in the html")
		return "", err
	}
	return qualityOptions, nil
}

func convertStringToClips(qualityOptionsTxt string) ([]clip, error) {
	clips := []clip{}
	startOfArray := strings.Index(qualityOptionsTxt, "[")
	endOfArray := strings.Index(qualityOptionsTxt, "]")
	if startOfArray == -1 || endOfArray == -1 {
		err := errors.New("invalid quailty options passed for conversion")
		return clips, err
	}
	jsonArrayTxt := qualityOptionsTxt[startOfArray : endOfArray+1]
	txtBytes := []byte(jsonArrayTxt)
	err := json.Unmarshal(txtBytes, &clips)
	return clips, err
}

func pickBestQuailty(clips []clip) (string, error) {
	clipSrc := ""

	for _, clip := range clips {
		if clip.Quality == "1080" || clip.Quality == "720" || clip.Quality == "360" {
			clipSrc = clip.Source
			break
		}
	}

	if clipSrc == "" {
		err := errors.New("clipSrc was not found in qualityOptions")
		return "", err
	}

	return clipSrc, nil
}
