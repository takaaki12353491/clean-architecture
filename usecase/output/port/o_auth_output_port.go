package outputport

import (
	outputdata "cln-arch/usecase/output/data"

	"golang.org/x/oauth2"
)

type OAuthOutputPort interface {
	Login(state string, url string) *outputdata.Login
	Callback(token string) *outputdata.Callback
	Auth(githubToken *oauth2.Token) *outputdata.Auth
}
