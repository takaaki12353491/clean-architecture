package controller

import (
	"cln-arch/config"
	"cln-arch/interface/gateway/database"
	"cln-arch/interface/presenter"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	"cln-arch/usecase/interactor"
	outputdata "cln-arch/usecase/output/data"
	"context"
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

func (ct *OAuthController) Login(session *inputdata.Session) (*outputdata.Login, error) {
	state := createRand()
	serverConf := config.NewServer()
	url := serverConf.Github.AuthCodeURL(state)
	expiry := time.Now().Add(10 * time.Minute)
	login := &inputdata.Login{
		State: state,
		URL:   url,
	}
	return ct.inputport.Login(login, session, &expiry)
}

func (ct *OAuthController) Callback(callback *inputdata.Callback) (*outputdata.Callback, error) {
	serverConf := config.NewServer()
	// make github token
	token, err := serverConf.Github.Exchange(context.Background(), callback.Code)
	if err != nil {
		return nil, err
	}
	githubToken := &inputdata.GithubToken{Token: token}
	// make user token
	expiry := time.Now().Add(7 * 24 * time.Hour)
	userToken := &inputdata.UserToken{
		Token:  createRand(),
		Expiry: &expiry,
	}
	return ct.inputport.Callback(callback, githubToken, userToken)
}

func (ct *OAuthController) Auth(_auth *inputdata.Auth) (*outputdata.Auth, error) {
	auth, err := ct.inputport.Auth(_auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}
