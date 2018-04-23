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

func Upload(fileStream io.ReadCloser, authString string, videoDescription string, channelName string) error {
	ytAuth, err := decodeAuth(authString)
	if err != nil {
		return err
	}
	authClient := getOAuthClient(ytAuth)
	uploadToYouTube(fileStream, authClient, videoDescription, channelName)
	return nil
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
