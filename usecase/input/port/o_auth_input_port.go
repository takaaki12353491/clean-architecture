package inputport

import (
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthInputPort interface {
	SetupGithubLogin(oauth *inputdata.OAuth) (*outputdata.Login, error)
}
