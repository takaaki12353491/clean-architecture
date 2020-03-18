package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"
)

type OAuthDatabase struct {
	SQLHandler
}

func NewOAuthDatabase() repository.OAuthRepository {
	return &OAuthDatabase{}
}

func (db *OAuthDatabase) FindUserStateBySessionID(string) (*model.UserState, error) {
	return nil, nil
}
func (db *OAuthDatabase) FindUserStateBySessionIDAndUserToken(sessionID string, token string) (*model.UserState, error) {
	return nil, nil
}
func (db *OAuthDatabase) FindOAuthTokenByUserID(string) (*model.OAuthToken, error) {
	return nil, nil
}
func (db *OAuthDatabase) StoreUserState(*model.UserState) error {
	return nil
}
func (db *OAuthDatabase) StoreOAuthToken(*model.OAuthToken) error {
	return nil
}
func (db *OAuthDatabase) StoreUserToken(*model.UserToken) error {
	return nil
}
