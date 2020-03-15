package inputdata

import (
	"time"

	"golang.org/x/oauth2"
)

type OAuth struct {
	ServerConf *ServerConf
	Session    *Session
	State      string
	URL        string
	Expiry     *time.Time
}

// DBConf is config using DB
type DBConf struct {
	Database string
	DSN      string
}

// ServerConf is above all
type ServerConf struct {
	DBConf DBConf
	Github oauth2.Config
}

// Session is recieved from server
type Session struct {
	ID string `json:"session_id"`
}

// Callback is callback param after github login
type Callback struct {
	Session
	Code  string `json:"code"`
	State string `json:"state"`
}

// Auth uses to authenticate user
type Auth struct {
	Session
	Token string `json:"token"`
}

// UserForOAuth is user's info
type UserForOAuth struct {
	Token string `json:"token"`
}

// GithubToken is github token
type GithubToken struct {
	oauth2.Token
}
