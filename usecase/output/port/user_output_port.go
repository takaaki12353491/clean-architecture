package outputport

import (
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
)

type UserOutputPort interface {
	View(*model.User) *outputdata.UserOutputData
}
