package inputdata

import (
	"cln-arch/config"
	"time"

	"golang.org/x/oauth2"
)

type Login struct {
	ServerConf *config.ServerConf
	Session    *Session
	State      string
	URL        string
	Expiry     *time.Time
}

// Callback is callback param after github login
type Callback struct {
	ServerConf  *config.ServerConf
	GithubToken *oauth2.Token
	Session
	UserToken *UserToken
	Code      string `json:"code"`
	State     string `json:"state"`
}

// Auth uses to authenticate user
type Auth struct {
	Session
	Token string `json:"token"`
}

// Session is recieved from server
type Session struct {
	ID string `json:"session_id"`
}

// UserToken
type UserToken struct {
	Token  string
	Expiry time.Time
}

// UserForOAuth is user's info
type UserForOAuth struct {
	Token string `json:"token"`
}
