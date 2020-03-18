package outputport

import (
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthOutputPort interface {
	Login(*model.Login) *outputdata.Login
	Callback(*model.UserToken) *outputdata.Callback
	Auth(*model.OAuthToken) *outputdata.Auth
}
