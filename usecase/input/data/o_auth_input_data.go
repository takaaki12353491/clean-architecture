package inputdata

import (
	"golang.org/x/oauth2"
)

// CallbackRequest is callback param after github login
type CallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type User struct {
	ID        uint
	Name      string
	AvatarURL string
}

type Callback struct {
	Code       string
	State      string
	User       *User
	OAuthToken *oauth2.Token
}
