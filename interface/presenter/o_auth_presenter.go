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

func (pre *OAuthPresenter) Auth(state *model.OAuthState) *outputdata.Auth {
	return &outputdata.Auth{
		State: state.State,
	}
}

func (pre *OAuthPresenter) Callback(user *model.User) *outputdata.Callback {
	return &outputdata.Callback{
		ID:        user.ID,
		Name:      user.Name,
		AvatorURL: user.AvatorURL,
	}
}
