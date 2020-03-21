package repository

import "cln-arch/domain/model"

type OAuthTokenRepository interface {
	Store(*model.OAuthToken) error
}
