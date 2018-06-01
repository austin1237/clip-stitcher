package uploader

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type YtAuth struct {
	ClientID     string    `json:"clientID"`
	ClientSecret string    `json:"clientSecret"`
	AccessToken  string    `json:"accessToken"`
	TokenType    string    `json:"tokenType"`
	RefreshToken string    `json:"refreshToken"`
	Expiry       time.Time `json:"expiry"`
}

type Video struct {
	FileStream       io.Reader
	VideoDescription string
	ChannelName      string
}

func Upload(video Video, pipeWriter *io.PipeWriter, authString string, done chan error) {
	defer pipeWriter.Close()
	ytAuth, err := decodeAuth(authString)
	if err != nil {
		done <- err
	}
	authClient := getOAuthClient(ytAuth)
	err = uploadToYouTube(video, authClient)
	if err != nil {
		done <- err
	}
	done <- nil
}

func decodeAuth(authString string) (YtAuth, error) {
	ytAuth := YtAuth{}
	decoded, err := base64.StdEncoding.DecodeString(authString)
	if err != nil {
		fmt.Println("decode error:", err)
		return ytAuth, err
	}

	err = json.Unmarshal([]byte(decoded), &ytAuth)
	if err != nil {
		fmt.Println("decode error:", err)
		return ytAuth, err
	}

	return ytAuth, nil

}
