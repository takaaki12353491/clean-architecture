package config

import (
	"cln-arch/consts"
	"cln-arch/errs"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	domain = "http://localhost:8080"
)

func OAuthConfig(service string) (*oauth2.Config, error) {
	switch service {
	case consts.Github:
		return githubConfig(), nil
	default:
		return nil, errs.Invalidated.New("invalid service")
	}
}

func githubConfig() *oauth2.Config {
	scopes := []string{"read:user"}
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  "http://localhost:8080/auth/github/callback",
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
	}
}

func twitterConfig() *oauth2.Config {
	scopes := []string{"read:user"}
	return &oauth2.Config{
		ClientID:     os.Getenv("TWITTER_CLIENT_ID"),
		ClientSecret: os.Getenv("TWITTER_CLIENT_SECRET"),
		RedirectURL:  "",
		Scopes:       scopes,
	}
}
