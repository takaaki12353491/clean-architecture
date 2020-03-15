package repository

import (
	"cln-arch/errs"
	"time"
)

// OAuthRepository is ...
type OAuthRepository interface {
	StoreState(state string, sessionID string, expiry *time.Time) errs.HTTPError
	FindBySessionID(string) (string, errs.HTTPError)
	FindBySessionIDAndUserToken(string, string) (*time.Time, int, errs.HTTPError)
	FindByUserTokenID(int) (string, string, string, *time.Time, errs.HTTPError)
	StoreUserToken(string, string, *time.Time, int) (int, errs.HTTPError)
	StoreGithubToken(string, string, string, *time.Time) (int, errs.HTTPError)
}
