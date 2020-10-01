package inputdata

import (
	"golang.org/x/oauth2"
)

// CallbackRequest is callback param after github login
type CallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type GithubUser struct {
	ID        uint
	Name      string
	AvatarURL string
}

type Callback struct {
	Code       string
	State      string
	User       *GithubUser
	OAuthToken *oauth2.Token
}
