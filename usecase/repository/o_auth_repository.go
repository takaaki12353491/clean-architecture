package repository

import (
	inputdata "cln-arch/usecase/input/data"
	"time"

	"golang.org/x/oauth2"
)

// OAuthRepository is ...
type OAuthRepository interface {
	StoreState(state string, sessionID string, expiry *time.Time) error
	FindBySessionID(string) (string, error)
	FindBySessionIDAndUserToken(string, string) (*time.Time, int, error)
	FindByUserTokenID(id int) (*oauth2.Token, error)
	StoreUserToken(sessionID string, userToken *inputdata.UserToken, id int) (int, error)
	StoreGithubToken(*oauth2.Token) (int, error)
}
