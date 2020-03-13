package port

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/output/data"
)

type UserOutputPort interface {
	View(*model.User) *data.UserOutputData
}
