package resfilter

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GetBestResoultionReturnsHigh(t *testing.T) {

	testClips := VideoResGroup{
		High:   []string{"high1", "high2"},
		Medium: []string{"medium1", "medium2"},
		Low:    []string{"low1", "low2"},
	}
	clips, err := getBestResoultion(testClips)
	assert.Nil(t, err)
	assert.Equal(t, clips, testClips.High)
}

func GetBestResoultionReturnsMedium(t *testing.T) {

	testClips := VideoResGroup{
		High:   []string{"high1"},
		Medium: []string{"medium1", "medium2"},
		Low:    []string{"low1", "low2"},
	}
	clips, err := getBestResoultion(testClips)
	assert.Nil(t, err)
	assert.Equal(t, clips, testClips.Medium)
}

func GetBestResoultionReturnsErr(t *testing.T) {
	expectedErr := errors.New("No clips found with high enough resoultion")
	testClips := VideoResGroup{
		High:   []string{"high1"},
		Medium: []string{"medium1"},
		Low:    []string{"low1", "low2"},
	}
	_, err := getBestResoultion(testClips)
	assert.Equal(t, err, expectedErr)

	testClips = VideoResGroup{
		High:   []string{},
		Medium: []string{},
		Low:    []string{},
	}
	_, err = getBestResoultion(testClips)
	assert.Equal(t, err, expectedErr)
}
