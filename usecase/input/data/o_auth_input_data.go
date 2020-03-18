package inputdata

import (
	"time"

	"golang.org/x/oauth2"
)

// Session is recieved from server
type Session struct {
	ID string `json:"session_id"`
}

// Login is auth login info
type Login struct {
	UserID  string
	State   string `json:"state"`
	URL     string `json:"redirect_url"`
	Session *Session
	Expiry  *time.Time
}

type Callback struct {
	Request    *CallbackRequest
	OAuthToken *oauth2.Token
	Token      string
	Expiry     *time.Time
}

// CallbackRequest is callback param after github login
type CallbackRequest struct {
	Session *Session
	Code    string `json:"code"`
	State   string `json:"state"`
}

// Auth uses to authenticate user
type Auth struct {
	Session *Session
	Token   string `json:"token"`
}

// User is user's info
type UserToken struct {
	Token  string `json:"token"`
	Expiry *time.Time
}
