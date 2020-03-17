package repository

import (
	"cln-arch/domain/model"
)

type UserRepository interface {
	FindByID(id string) (*model.User, error)
	Create(user *model.User) error
	Update(user *model.User) error
	Delete(id string) error
}
