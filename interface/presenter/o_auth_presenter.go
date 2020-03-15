package presenter

import (
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
