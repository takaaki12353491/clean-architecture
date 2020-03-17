package presenter

import (
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
	outputport "cln-arch/usecase/output/port"
)

type OAuthPresenter struct {
}

func NewOAuthPresenter() outputport.OAuthOutputPort {
	return &OAuthPresenter{}
}

func (pre *OAuthPresenter) Login(state string, url string) *outputdata.Login {
	return &outputdata.Login{
		State: state,
		URL:   url,
	}
}

func (pre *OAuthPresenter) Callback(userToken *model.UserToken) *outputdata.Callback {
	return &outputdata.Callback{
		Token: userToken.Token,
	}
}

func (pre *OAuthPresenter) Auth(githubToken *model.GithubToken) *outputdata.Auth {
	return &outputdata.Auth{
		GithubToken: githubToken.Token,
	}
}
