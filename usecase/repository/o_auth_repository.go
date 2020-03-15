package repository

import (
	"time"
)

// OAuthRepository is ...
type OAuthRepository interface {
	StoreState(state string, sessionID string, expiry *time.Time) error
	FindBySessionID(string) (string, error)
	FindBySessionIDAndUserToken(string, string) (*time.Time, int, error)
	FindByUserTokenID(int) (string, string, string, *time.Time, error)
	StoreUserToken(string, string, *time.Time, int) (int, error)
	StoreGithubToken(string, string, string, *time.Time) (int, error)
}
