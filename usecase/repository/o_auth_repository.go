package repository

import (
	"cln-arch/domain/model"
	"time"
)

// OAuthRepository is ...
type OAuthRepository interface {
	FindStateBySession(session *model.Session) (string, error)
	FindBySessionIDAndUserToken(sessionID string, token string) (*time.Time, int, error)
	FindByUserTokenID(id int) (*model.GithubToken, error)
	StoreState(state string, session *model.Session, expiry *time.Time) error
	StoreUserToken(model *model.Session, userToken *model.UserToken, id int) (int, error)
	StoreGithubToken(githubToken *model.GithubToken) (int, error)
}
