package inputport

import (
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthInputPort interface {
	Login(oauth *inputdata.Login) (*outputdata.Login, error)
	Callback(callback *inputdata.Callback) (*outputdata.Callback, error)
}
