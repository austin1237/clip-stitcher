package twitch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTimeFromURLFromHour(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=2h11m40s"
	cTime, err := getTimeFromURL(testURL)
	assert.Nil(t, err)
	assert.Equal(t, 2, cTime.Hours)
	assert.Equal(t, 11, cTime.Minutes)
	assert.Equal(t, 40, cTime.Seconds)
}

func TestGetTimeFromURLFromMinute(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=11m40s"
	cTime, err := getTimeFromURL(testURL)
	assert.Nil(t, err)
	assert.Equal(t, 0, cTime.Hours)
	assert.Equal(t, 11, cTime.Minutes)
	assert.Equal(t, 40, cTime.Seconds)
}

func TestGetTimeFromURLFromSecond(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=40s"
	cTime, err := getTimeFromURL(testURL)
	assert.Nil(t, err)
	assert.Equal(t, 0, cTime.Hours)
	assert.Equal(t, 0, cTime.Minutes)
	assert.Equal(t, 40, cTime.Seconds)
}

func TestInvalidHoursFromURL(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=Ah11m40s"
	_, err := getTimeFromURL(testURL)
	assert.Equal(t, "invalid hours found in vod url", err.Error())
}

func TestInvalidMinutesFromURL(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=2hJm40s"
	_, err := getTimeFromURL(testURL)
	assert.Equal(t, "invalid minutes found in vod url", err.Error())
}

func TestInvalidSecondsFromURL(t *testing.T) {
	testURL := "https://www.twitch.tv/videos/240497733?t=1h11mbs"
	_, err := getTimeFromURL(testURL)
	assert.Equal(t, "invalid seconds found in vod url", err.Error())
}
