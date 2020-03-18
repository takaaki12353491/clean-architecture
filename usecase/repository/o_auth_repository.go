package repository

import (
	"cln-arch/domain/model"
)

// OAuthRepository is ...
type OAuthRepository interface {
	FindUserStateBySessionID(string) (*model.UserState, error)
	FindUserStateBySessionIDAndUserToken(sessionID string, token string) (*model.UserState, error)
	FindOAuthTokenByUserID(string) (*model.OAuthToken, error)
	StoreUserState(*model.UserState) error
	StoreOAuthToken(*model.OAuthToken) error
	StoreUserToken(*model.UserToken) error
}
