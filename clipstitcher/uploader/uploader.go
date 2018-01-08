package uploader

import "time"

func Upload(clipLocation string, ytToken string, ytSecret string, ytAccess string, ytRefresh string, ytExpiriy time.Time) {
	authClient := getOAuthClient(ytToken, ytSecret, ytAccess, ytRefresh, ytExpiriy)
	uploadToYouTube(clipLocation, authClient)
}
