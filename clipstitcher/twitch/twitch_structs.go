package twitch

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
	} `json:"clips"`
}

type PreparedClips struct {
	VideoLinks       []string
	VideoDescription string
}

type asyncString struct {
	value string
	err   error
}
