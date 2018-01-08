package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scrapVidSrcOnPage(clipURL string, response chan asyncString) {
	doc, err := goquery.NewDocument(clipURL)
	javascript, err := findJavaScript(doc)
	if err != nil {
		response <- asyncString{"", err}
		return
	}
	qualityOptionsTxt, err := findTheQualtyOptions(javascript)

	if err != nil {
		response <- asyncString{"", err}
		return
	}
	clipSrc, err := findClpSrc(qualityOptionsTxt)

	if err != nil {
		response <- asyncString{"", err}
		return
	}
	fmt.Println("Clip src is " + clipSrc)
	response <- asyncString{clipSrc, nil}

}

func findJavaScript(doc *goquery.Document) (string, error) {
	scriptText := ""
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptText = s.Text()
		if strings.Contains(scriptText, "quality_options") {
			scriptText = strings.Replace(scriptText, " ", "", -1)
		}
	})

	if scriptText == "" {
		err := errors.New("no javascript with clip src found on the page")
		return "", err
	}

	return scriptText, nil
}

func findTheQualtyOptions(javaScriptText string) (string, error) {
	qualityOptions := ""
	// Loops through the multiline js script
	for _, line := range strings.Split(strings.TrimSuffix(javaScriptText, "\n"), "\n") {
		if strings.Contains(line, "quality_options") {
			qualityOptions = line
			break
		}
	}
	if qualityOptions == "" {
		err := errors.New("quality_options not found in the javascript")
		return "", err
	}
	return qualityOptions, nil
}

func convertStringToClips(qualityOptionsTxt string) ([]clip, error) {
	clips := []clip{}
	startOfArray := strings.Index(qualityOptionsTxt, "[")
	endOfArray := strings.Index(qualityOptionsTxt, "]")
	jsonArrayTxt := qualityOptionsTxt[startOfArray : endOfArray+1]
	txtBytes := []byte(jsonArrayTxt)
	err := json.Unmarshal(txtBytes, &clips)
	return clips, err
}

func findClpSrc(qualityOptionsTxt string) (string, error) {
	clipSrc := ""
	clips, err := convertStringToClips(qualityOptionsTxt)
	if err != nil {
		return "", err
	}

	for _, clip := range clips {
		if clip.Quality == "1080" || clip.Quality == "720" || clip.Quality == "360" {
			clipSrc = clip.Source
			break
		}
	}

	return clipSrc, nil
}
