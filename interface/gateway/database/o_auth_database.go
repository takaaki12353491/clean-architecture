package database

import (
	"cln-arch/errs"
	"cln-arch/usecase/repository"
	"time"
)

type OAuthDatabase struct {
}

func NewOAuthDatabase() repository.OAuthRepository {
	return &OAuthDatabase{}
}

func (db *OAuthDatabase) StoreState(state string, sessionID string, expiry *time.Time) errs.HTTPError {
	return nil
}

func (db *OAuthDatabase) FindBySessionID(string) (string, errs.HTTPError) {
	return "", nil
}

func (db *OAuthDatabase) FindBySessionIDAndUserToken(string, string) (*time.Time, int, errs.HTTPError) {
	return nil, 0, nil
}

func (db *OAuthDatabase) FindByUserTokenID(int) (string, string, string, *time.Time, errs.HTTPError) {
	return "", "", "", nil, nil
}

func (db *OAuthDatabase) StoreUserToken(string, string, *time.Time, int) (int, errs.HTTPError) {
	return 0, nil
}

func (db *OAuthDatabase) StoreGithubToken(string, string, string, *time.Time) (int, errs.HTTPError) {
	return 0, nil
}
