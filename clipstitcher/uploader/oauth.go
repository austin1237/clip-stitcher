package uploader

import (
	"net/http"

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
func getOAuthClient(ytAuth YtAuth) *http.Client {
	ctx := context.Background()
	config := generateConfig(ytAuth.ClientID, ytAuth.ClientSecret)
	tok := &oauth2.Token{
		AccessToken:  ytAuth.AccessToken,
		RefreshToken: ytAuth.RefreshToken,
		TokenType:    "Bearer",
		Expiry:       ytAuth.Expiry,
	}

	return config.Client(ctx, tok)
}
