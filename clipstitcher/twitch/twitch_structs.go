package twitch

type clip struct {
	Quality   string  `json:"quality"`
	Source    string  `json:"source"`
	FrameRate float32 `json:"frame_rate"`
}

type twitchResponseUrls struct {
	Clips []struct {
		URL string `json:"url"`
	} `json:"clips"`
}

type asyncString struct {
	value string
	err   error
}
