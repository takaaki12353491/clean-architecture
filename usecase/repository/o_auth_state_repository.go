package repository

import "cln-arch/domain/model"

type OAuthStateRepository interface {
	FindByState(string) (*model.OAuthState, error)
	Store(*model.OAuthState) error
	Delete(*model.OAuthState) error
}
