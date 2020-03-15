package database

import (
	"time"

	"golang.org/x/oauth2"

	inputdata "cln-arch/usecase/input/data"
	"cln-arch/usecase/repository"
)

type OAuthDatabase struct {
}

func NewOAuthDatabase() repository.OAuthRepository {
	return &OAuthDatabase{}
}

func (db *OAuthDatabase) StoreState(state string, sessionID string, expiry *time.Time) error {
	return nil
}

func (db *OAuthDatabase) FindBySessionID(string) (string, error) {
	return "", nil
}

func (db *OAuthDatabase) FindBySessionIDAndUserToken(string, string) (*time.Time, int, error) {
	return nil, 0, nil
}

func (db *OAuthDatabase) FindByUserTokenID(int) (string, string, string, *time.Time, error) {
	return "", "", "", nil, nil
}

func (db *OAuthDatabase) StoreUserToken(sessionID string, userToken *inputdata.UserToken, id int) (int, error) {
	return 0, nil
}

func (db *OAuthDatabase) StoreGithubToken(githubToken *oauth2.Token) (int, error) {
	return 0, nil
}
