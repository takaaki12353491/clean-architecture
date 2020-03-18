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

func (pre *OAuthPresenter) Login(login *model.Login) *outputdata.Login {
	return &outputdata.Login{
		State: login.State,
		URL:   login.URL,
	}
}

func (pre *OAuthPresenter) Callback(userToken *model.UserToken) *outputdata.Callback {
	return &outputdata.Callback{
		Token: userToken.Token,
	}
}

func (pre *OAuthPresenter) Auth(oauthToken *model.OAuthToken) *outputdata.Auth {
	return &outputdata.Auth{
		OAuthToken: oauthToken.Token,
	}
}
