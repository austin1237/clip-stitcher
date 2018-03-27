package twitch

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var apiEnpoint = "https://api.twitch.tv/kraken/clips/"

func twitchClipRequest(streamName string, desiredCount int, clientID string) *http.Request {
	queryParams := "top?period=day&channel=" + streamName + "&limit=" + strconv.Itoa(desiredCount)
	url := apiEnpoint + queryParams
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/vnd.twitchtv.v5+json")
	req.Header.Set("Client-ID", clientID)
	return req
}

func unMarshalClipJSON(resp *http.Response) (twitchClips, error) {
	tclips := twitchClips{}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &tclips)
	if err != nil {
		fmt.Println(err.Error())
		err = errors.New("an error occured trying to parse twitch's json")
		return twitchClips{}, err
	}
	return tclips, nil
}

func getClips(streamName string, desiredCount int, clientID string) (twitchClips, error) {
	client := &http.Client{}
	req := twitchClipRequest(streamName, desiredCount, clientID)
	resp, err := client.Do(req)
	if err != nil {
		err = errors.New("an error occured trying connect to twitch's clip api")
		return twitchClips{}, err
	}
	clips, err := unMarshalClipJSON(resp)
	if err != nil {
		err = errors.New("an error occured trying connect to twitch's clip api")
		return twitchClips{}, err
	}
	return clips, nil
}
