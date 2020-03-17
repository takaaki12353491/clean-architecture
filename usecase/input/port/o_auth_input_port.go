package inputport

import (
	inputdata "cln-arch/usecase/input/data"
	outputdata "cln-arch/usecase/output/data"
	"time"
)

type OAuthInputPort interface {
	Login(login *inputdata.Login, _session *inputdata.Session, expiry *time.Time) (*outputdata.Login, error)
	Callback(callback *inputdata.Callback, _githubToken *inputdata.GithubToken, _userToken *inputdata.UserToken) (*outputdata.Callback, error)
	Auth(auth *inputdata.Auth) (*outputdata.Auth, error)
}
