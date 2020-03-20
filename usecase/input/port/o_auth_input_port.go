package inputport

import (
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthInputPort interface {
	Auth() (*outputdata.Auth, error)
	Callback(*inputdata.Callback) (*outputdata.Callback, error)
}
