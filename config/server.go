package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func NewGithubConf() *oauth2.Config {
	scopes := []string{"read:user"}
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/github/callback",
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
	}
}
