package repository

import (
	"cln-arch/errs"
	inputdata "cln-arch/usecase/input/data"
)

type UserRepository interface {
	Create(inputdata.User) errs.HTTPError
}
