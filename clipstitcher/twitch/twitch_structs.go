package twitch

import "time"

type clip struct {
	Quality   string  `json:"quality"`
	Source    string  `json:"source"`
	FrameRate float32 `json:"frame_rate"`
}

type twitchClips struct {
	Clips []struct {
		URL      string  `json:"url"`
		Duration float64 `json:"duration"`
		Title    string  `json:"title"`
		Vod      struct {
			ID  int    `json:"id"`
			URL string `json:"url"`
		} `json:"vod"`
	} `json:"clips"`
}

type clipDuration struct {
	StartTime time.Time
	EndTime   time.Time
}

type clipTime struct {
	Hours   int
	Minutes int
	Seconds int
}

type PreparedClips struct {
	VideoLinks       []string
	VideoDescription string
}

type asyncString struct {
	value string
	err   error
}
