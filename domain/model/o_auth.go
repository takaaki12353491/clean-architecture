package model

import (
	"cln-arch/validator"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type OAuthState struct {
	State  string
	Expiry *time.Time
}

func NewOAuthState() (*OAuthState, error) {
	state := createRand()
	expiry := time.Now().Add(10 * time.Minute)
	oauthState := &OAuthState{
		State:  state,
		Expiry: &expiry,
	}
	err := validator.Validate(oauthState)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return oauthState, nil
}

// OAuthToken is oauth token
type OAuthToken struct {
	User  *User
	Token *oauth2.Token
}

func NewOAuthToken(user *User, token *oauth2.Token) (*OAuthToken, error) {
	oauthToken := &OAuthToken{
		User:  user,
		Token: token,
	}
	err := validator.Validate(oauthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return oauthToken, nil
}
