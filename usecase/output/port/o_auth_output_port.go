package outputport

import (
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
)

type OAuthOutputPort interface {
	Login(state string, url string) *outputdata.Login
	Callback(userToken *model.UserToken) *outputdata.Callback
	Auth(githubToken *model.GithubToken) *outputdata.Auth
}
