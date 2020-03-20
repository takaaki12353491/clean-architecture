package inputdata

import (
	"golang.org/x/oauth2"
)

type Callback struct {
	Request    *CallbackRequest
	OAuthToken *oauth2.Token
}

// CallbackRequest is callback param after github login
type CallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}
