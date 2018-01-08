package uploader

import (
	"log"
	"net/http"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func generateConfig(clientID string, clientSecret string) *oauth2.Config {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "oob",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}

	return config
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getOAuthClient(clientID string, clientSecret string, accessToken string, refreshToken string, expiry time.Time) *http.Client {
	ctx := context.Background()
	config := generateConfig(clientID, clientSecret)
	tok := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Expiry:       expiry,
	}

	return config.Client(ctx, tok)
}

// Exchange the authorization code for an access token
func exchangeToken(config *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := config.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token %v", err)
	}
	return tok, nil
}
