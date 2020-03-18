package model

import (
	"cln-arch/validator"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type Login struct {
	State string
	URL   string
}

func NewLogin(state string, url string) (*Login, error) {
	login := &Login{
		State: state,
		URL:   url,
	}
	err := validator.Validate(login)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return login, nil
}

type UserState struct {
	UserID    string
	SessionID string
	State     string
	Expiry    *time.Time
}

func NewUserState(userID string, sessionID string, state string, expiry *time.Time) (*UserState, error) {
	userState := &UserState{
		UserID:    userID,
		SessionID: sessionID,
		State:     state,
		Expiry:    expiry,
	}
	err := validator.Validate(userState)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return userState, nil
}

type UserToken struct {
	UserID string
	Token  string
	Expiry *time.Time
}

func NewUserToken(userID string, token string, expiry *time.Time) (*UserToken, error) {
	userToken := &UserToken{
		UserID: userID,
		Token:  token,
		Expiry: expiry,
	}
	err := validator.Validate(userToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return userToken, nil
}

// OAuthToken is oauth token
type OAuthToken struct {
	UserID string
	Token  *oauth2.Token
}

func NewOAuthToken(userID string, token *oauth2.Token) (*OAuthToken, error) {
	oauthToken := &OAuthToken{
		UserID: userID,
		Token:  token,
	}
	err := validator.Validate(oauthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return oauthToken, nil
}
