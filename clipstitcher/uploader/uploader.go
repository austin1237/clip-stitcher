package uploader

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"time"

	"github.com/pkg/errors"
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

func Upload(video Video, authString string) error {
	ytAuth, err := decodeAuth(authString)
	if err != nil {
		err = errors.New(err.Error())
		return err
	}
	authClient := getOAuthClient(ytAuth)
	err = uploadToYouTube(video, authClient)
	if err != nil {
		err = errors.New(err.Error())
		return err
	}
	return nil
}

func decodeAuth(authString string) (YtAuth, error) {
	ytAuth := YtAuth{}
	decoded, err := base64.StdEncoding.DecodeString(authString)
	if err != nil {
		err = errors.New(err.Error())
		return ytAuth, err
	}

	err = json.Unmarshal([]byte(decoded), &ytAuth)
	if err != nil {
		err = errors.New(err.Error())
		return ytAuth, err
	}

	return ytAuth, nil

}
