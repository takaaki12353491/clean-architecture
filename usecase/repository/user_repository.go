package repository

import (
	"cln-arch/domain/model"
	"cln-arch/errs"
	inputdata "cln-arch/usecase/input/data"
)

type UserRepository interface {
	FindByID(id string) (model.User, errs.HTTPError)
	Create(inputdata.User) errs.HTTPError
	Update(inputdata.User) errs.HTTPError
	Delete(id string) errs.HTTPError
}
