package config

import (
	"os"

	"golang.org/x/oauth2"
	oauth2github "golang.org/x/oauth2/github"
)

// ServerConf is above all
type ServerConf struct {
	DBConf *DBConf
	Github *oauth2.Config
}

// DBConf is config using DB
type DBConf struct {
	Database string
	DSN      string
}

func NewServer() *ServerConf {
	return &ServerConf{
		DBConf: getDBConf(),
		Github: getGithubConf(),
	}
}

func getDBConf() *DBConf {
	return &DBConf{
		Database: os.Getenv("DATABASE"),
		DSN:      os.Getenv("DSN"),
	}
}

func getGithubConf() *oauth2.Config {
	scopes := []string{"repo"}
	return &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SERVER_HOST"),
		Scopes:       scopes,
		Endpoint:     oauth2github.Endpoint,
	}
}
