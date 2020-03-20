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

func (db *OAuthDatabase) FindStateByState(state string) (*model.OAuthState, error) {
	oauthState := &model.OAuthState{}
	db.First(oauthState, "state = ?", state)
	return oauthState, nil
}

func (db *OAuthDatabase) FindTokenByToken(token string) (*model.OAuthToken, error) {
	oauthToken := &model.OAuthToken{}
	db.First(oauthToken, "access_token = ?", token)
	return oauthToken, nil
}

func (db *OAuthDatabase) StoreState(state *model.OAuthState) error {
	db.Create(state)
	return nil
}

func (db *OAuthDatabase) StoreToken(token *model.OAuthToken) error {
	db.Create(token)
	return nil
}
