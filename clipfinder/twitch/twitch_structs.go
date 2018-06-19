package twitch

import "time"

type clip struct {
	Quality   string  `json:"quality"`
	Source    string  `json:"source"`
	FrameRate float32 `json:"frame_rate"`
}

type twitchClip struct {
	Slug     string  `json:"slug"`
	URL      string  `json:"url"`
	Duration float64 `json:"duration"`
	Title    string  `json:"title"`
	Vod      struct {
		ID  string `json:"id"`
		URL string `json:"url"`
	} `json:"vod"`
}

type twitchAPIResp struct {
	Clips []twitchClip `json:"clips"`
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

type twitchDuration struct {
	URL       string
	VodID     string
	StartTime time.Time
	EndTime   time.Time
}

type PreparedClips struct {
	VideoSlugs       []string
	VideoDescription string
}

type asyncString struct {
	value string
	err   error
}
