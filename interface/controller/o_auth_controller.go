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
		ServerConf: serverConf,
		Session:    session,
		State:      state,
		URL:        url,
		Expiry:     &expiry,
	}
	return ct.inputport.Login(login)
}

func (ct *OAuthController) Callback(ctx context.Context, sessionID string, code string, state string) (*outputdata.Callback, error) {
	serverConf := config.NewServer()
	// make github token
	githubToken, err := serverConf.Github.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	session := inputdata.Session{
		ID: sessionID,
	}
	// make user token
	userToken := &inputdata.UserToken{
		Token:  createRand(),
		Expiry: time.Now().Add(7 * 24 * time.Hour),
	}
	callback := &inputdata.Callback{
		ServerConf:  serverConf,
		GithubToken: githubToken,
		Session:     session,
		UserToken:   userToken,
		Code:        code,
		State:       state,
	}
	return ct.inputport.Callback(callback)
}
