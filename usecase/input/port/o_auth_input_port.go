package inputport

import (
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthInputPort interface {
	Login(*inputdata.Login) (*outputdata.Login, error)
	Callback(*inputdata.Callback) (*outputdata.Callback, error)
	Auth(auth *inputdata.Auth) (*outputdata.Auth, error)
}
