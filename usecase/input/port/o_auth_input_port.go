package inputport

import (
	"cln-arch/errs"
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthInputPort interface {
	SetupGithubLogin(oauth *inputdata.OAuth) (*outputdata.Login, errs.HTTPError)
}
