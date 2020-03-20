package outputport

import (
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthOutputPort interface {
	Auth(*model.OAuthState) *outputdata.Auth
	Callback(*model.OAuthToken) *outputdata.Callback
}
