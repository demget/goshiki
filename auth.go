package goshiki

import (
	"context"

	"golang.org/x/oauth2"
)

// Auth returns a new AuthConfig with given data and oauth2 endpoints.
func Auth(id, secret, uri string) AuthConfig {
	return AuthConfig{oauth2.Config{
		ClientID:     id,
		ClientSecret: secret,
		RedirectURL:  uri,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://shikimori.org/oauth/authorize",
			TokenURL: "https://shikimori.org/oauth/token",
		},
	}}
}

// AuthConfig "overrides" a oauth2.Config to make interaction with it simpler.
// Use goshiki.Auth() to create new instance of config.
//
// Example:
//		auth := goshiki.Auth(id, secret, uri)
//
//		url := auth.URL() // authorize URL for user to get code
//		code := "..."     // pushed code to RedirectURI
//
//		token, err := auth.Token(code)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		api := goshiki.New("goshiki", token) // API with valid AccessToken
//
type AuthConfig struct {
	conf oauth2.Config
}

// URL returns an AuthCodeURL from oauth2.Config.
func (auth AuthConfig) URL() string {
	return auth.conf.AuthCodeURL("")
}

// Token returns Exchanged token from oauth2.Config with default context.
func (auth AuthConfig) Token(code string) (*oauth2.Token, error) {
	return auth.conf.Exchange(context.Background(), code)
}
