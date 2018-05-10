package twitch

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const mockOptions string = `quality_options:[{"quality":"1080","source":"https://clips-media-assets.twitch.tv/172801092.mp4","frame_rate":60},{"quality":"720","source":"https://clips-media-assets.twitch.tv/AT-172801092-1280x720.mp4","frame_rate":60},{"quality":"480","source":"https://clips-media-assets.twitch.tv/AT-172801092-854x480.mp4","frame_rate":30},{"quality":"360","source":"https://clips-media-assets.twitch.tv/AT-172801092-640x360.mp4","frame_rate":30}],`

func TestAsyncGetClipHTML(t *testing.T) {
	testURL := "https://clips.twitch.tv/AliveMistySoybeanDatSheffy"
	scrapResponse := make(chan asyncString)
	go asyncGetClipHTML(testURL, scrapResponse)
	response := <-scrapResponse
	assert.Nil(t, response.err)
	assert.NotEqual(t, "", response.value)
}

func TestFindQualityOptionsGreen(t *testing.T) {
	bytes, err := ioutil.ReadFile("test.html")
	testHTML := string(bytes)
	assert.Nil(t, err)
	actualOptions, err := findQualtyOptions(testHTML)
	assert.Nil(t, err)
	assert.Equal(t, mockOptions, actualOptions)
}

func TestFindQualityOptionsRed(t *testing.T) {
	bytes, err := ioutil.ReadFile("test.html")
	testHTML := string(bytes)
	assert.Nil(t, err)
	actualOptions, err := findQualtyOptions(testHTML)
	assert.Nil(t, err)
	assert.Equal(t, mockOptions, actualOptions)
}

func TestConvertStringToClipsGreen(t *testing.T) {
	clips, err := convertStringToClips(mockOptions)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(clips))
}

func TestConvertStringToClipsRed(t *testing.T) {
	_, err := convertStringToClips("")
	assert.NotNil(t, err)

	_, err = convertStringToClips("i[nvalidjso]n")
	assert.NotNil(t, err)
}

func TestPickBestQualityGreen(t *testing.T) {
	mockClips := []clip{
		clip{
			Quality:   "1080",
			Source:    "1080clip.mp4",
			FrameRate: 60,
		},
		clip{
			Quality:   "720",
			Source:    "720clip.mp4",
			FrameRate: 60,
		},
		clip{
			Quality:   "360",
			Source:    "360clip.mp4",
			FrameRate: 60,
		},
	}
	clipSrc, err := pickBestQuailty(mockClips)
	assert.Nil(t, err)
	assert.Equal(t, "1080clip.mp4", clipSrc)
}

func TestPickBestQualityRed(t *testing.T) {
	mockClips := []clip{}
	_, err := pickBestQuailty(mockClips)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "clipSrc was not found in qualityOptions")
}
