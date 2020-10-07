package config

import (
	"cln-arch/consts"
	"cln-arch/errs"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	domain    = "http://192.168.50.10:8080"
	directory = "/oauth/callback"
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
		RedirectURL:  strings.Join([]string{domain, directory, "?service=github"}, ""),
		Scopes:       scopes,
		Endpoint:     github.Endpoint,
	}
}
