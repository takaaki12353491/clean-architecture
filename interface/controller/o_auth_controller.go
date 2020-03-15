package controller

import (
	"cln-arch/interface/gateway/database"
	"cln-arch/interface/presenter"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	"cln-arch/usecase/interactor"
	outputdata "cln-arch/usecase/output/data"
	"time"
)

type OAuthController struct {
	inputport inputport.OAuthInputPort
}

func NewOAuthController() *OAuthController {
	return &OAuthController{
		inputport: interactor.NewOAuthInteractor(
			database.NewOAuthDatabase(),
			presenter.NewOAuthPresenter(),
		),
	}
}

func (ct *OAuthController) Login(conf *inputdata.ServerConf, session *inputdata.Session) (*outputdata.Login, error) {
	state := createRand()
	url := conf.Github.AuthCodeURL(state)
	expiry := time.Now().Add(10 * time.Minute)
	oauth := &inputdata.OAuth{
		ServerConf: conf,
		Session:    session,
		State:      state,
		URL:        url,
		Expiry:     &expiry,
	}
	return ct.inputport.SetupGithubLogin(oauth)
}
