package uploader

import (
	"io"
	"time"
)

func Upload(fileStream io.ReadCloser, ytToken string, ytSecret string, ytAccess string, ytRefresh string, ytExpiriy time.Time) {
	authClient := getOAuthClient(ytToken, ytSecret, ytAccess, ytRefresh, ytExpiriy)
	uploadToYouTube(fileStream, authClient)
}
