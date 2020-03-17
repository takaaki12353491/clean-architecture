package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"
	"time"
)

type OAuthDatabase struct {
	SQLHandler
}

func NewOAuthDatabase() repository.OAuthRepository {
	return &OAuthDatabase{}
}

func (db *OAuthDatabase) FindStateBySession(session *model.Session) (string, error) {
	return "", nil
}
func (db *OAuthDatabase) FindBySessionIDAndUserToken(sessionID string, token string) (*time.Time, int, error) {
	return nil, 0, nil
}
func (db *OAuthDatabase) FindByUserTokenID(id int) (*model.GithubToken, error) {
	return &model.GithubToken{}, nil
}
func (db *OAuthDatabase) StoreState(state string, session *model.Session, expiry *time.Time) error {
	return nil
}
func (db *OAuthDatabase) StoreUserToken(model *model.Session, userToken *model.UserToken, id int) (int, error) {
	return 0, nil
}
func (db *OAuthDatabase) StoreGithubToken(githubToken *model.GithubToken) (int, error) {
	return 0, nil
}
