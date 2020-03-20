package repository

import (
	"cln-arch/domain/model"
)

// OAuthRepository is ...
type OAuthRepository interface {
	FindStateByState(string) (*model.OAuthState, error)
	FindTokenByToken(string) (*model.OAuthToken, error)
	StoreState(*model.OAuthState) error
	StoreToken(*model.OAuthToken) error
}
