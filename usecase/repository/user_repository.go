package repository

import (
	"cln-arch/domain/model"
)

type UserRepository interface {
	FindByID(uint) (*model.User, error)
	Store(*model.User) error
	Update(*model.User) error
}
