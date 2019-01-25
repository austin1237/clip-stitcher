package resfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBestResoultionReturnsHigh(t *testing.T) {

	testClips := VideoResGroup{
		High:   []string{"high1", "high2"},
		Medium: []string{"medium1", "medium2"},
		Low:    []string{"low1", "low2"},
	}
	clips, err := getBestResoultion(testClips)
	assert.Nil(t, err)
	assert.Equal(t, clips, testClips.High)
}

func TestBestResoultionReturnsMedium(t *testing.T) {

	testClips := VideoResGroup{
		High:   []string{"high1"},
		Medium: []string{"medium1", "medium2"},
		Low:    []string{"low1", "low2"},
	}
	clips, err := getBestResoultion(testClips)
	assert.Nil(t, err)
	assert.Equal(t, clips, testClips.Medium)
}

func TestBestResoultionReturnsSuperMedium(t *testing.T) {

	testClips := VideoResGroup{
		High:        []string{"high1"},
		SuperMedium: []string{"medium1", "medium2"},
		Medium:      []string{"medium1"},
		Low:         []string{"low1", "low2"},
	}
	clips, err := getBestResoultion(testClips)
	assert.Nil(t, err)
	assert.Equal(t, clips, testClips.SuperMedium)
}

func TestBestResoultionReturnsErr(t *testing.T) {
	testClips := VideoResGroup{
		High:   []string{"high1"},
		Medium: []string{"medium1"},
		Low:    []string{"low1", "low2"},
	}
	_, err := getBestResoultion(testClips)
	assert.Equal(t, ErrNoHighResoultionClip, err)

	testClips = VideoResGroup{
		High:   []string{},
		Medium: []string{},
		Low:    []string{},
	}
	_, err = getBestResoultion(testClips)
	assert.Equal(t, ErrNoHighResoultionClip, err)
}
