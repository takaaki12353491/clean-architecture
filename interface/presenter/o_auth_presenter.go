package presenter

import (
	"cln-arch/config"
	"cln-arch/domain/model"
	outputdata "cln-arch/usecase/output/data"
	outputport "cln-arch/usecase/output/port"
)

type OAuthPresenter struct {
}

func NewOAuthPresenter() outputport.OAuthOutputPort {
	return &OAuthPresenter{}
}

func (pre *OAuthPresenter) Auth(state *model.OAuthState) *outputdata.Auth {
	oauthConfig := config.NewGithubConf()
	url := oauthConfig.AuthCodeURL(state.State)
	return &outputdata.Auth{
		URL: url,
	}
}

func (pre *OAuthPresenter) Callback(user *model.User) *outputdata.Callback {
	return &outputdata.Callback{
		ID:        user.ID,
		Name:      user.Name,
		AvatorURL: user.AvatorURL,
	}
}
