package model

import (
	"cln-arch/validator"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type OAuthState struct {
	gorm.Model
	State  string    `gorm:"column:state"`
	Expiry time.Time `gorm:"column:expiry"`
}

func NewOAuthState() (*OAuthState, error) {
	oauthState := new(OAuthState)
	oauthState.ID = uint(uuid.New().ID())
	oauthState.State = createRand()
	oauthState.Expiry = time.Now().Add(10 * time.Minute)
	err := validator.Validate(oauthState)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return oauthState, nil
}

// OAuthToken is oauth token
type OAuthToken struct {
	User         User
	AccessToken  string    `gorm:"column:access_token"`
	TokenType    string    `gorm:"column:token_type"`
	RefreshToken string    `gorm:"column:refresh_token"`
	Expiry       time.Time `gorm:"column:expiry"`
}

func NewOAuthToken(user *User, token *oauth2.Token) (*OAuthToken, error) {
	oauthToken := &OAuthToken{
		User:         *user,
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	err := validator.Validate(oauthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return oauthToken, nil
}
