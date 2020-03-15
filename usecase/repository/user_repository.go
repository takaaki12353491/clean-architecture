package repository

import (
	"cln-arch/domain/model"
	inputdata "cln-arch/usecase/input/data"
)

type UserRepository interface {
	FindByID(id string) (model.User, error)
	Create(inputdata.User) error
	Update(inputdata.User) error
	Delete(id string) error
}
