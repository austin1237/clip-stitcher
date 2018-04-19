package uploader

import (
	"io"
	"time"
)

func Upload(fileStream io.ReadCloser, ytToken string, ytSecret string, ytAccess string, ytRefresh string, ytExpiriy time.Time, videoDescription string, channelName string) {
	authClient := getOAuthClient(ytToken, ytSecret, ytAccess, ytRefresh, ytExpiriy)
	uploadToYouTube(fileStream, authClient, videoDescription, channelName)
}
