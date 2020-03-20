package controller

import (
	"cln-arch/config"
	"cln-arch/infra/database"
	"cln-arch/interface/presenter"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	"cln-arch/usecase/interactor"
	outputdata "cln-arch/usecase/output/data"
	"context"

	log "github.com/sirupsen/logrus"
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

func (ct *OAuthController) Auth() (*outputdata.Auth, error) {
	return ct.inputport.Auth()
}

func (ct *OAuthController) Callback(iCallbackRequest *inputdata.CallbackRequest) (*outputdata.Callback, error) {
	githubConf := config.NewGithubConf()
	oauthToken, err := githubConf.Exchange(context.Background(), iCallbackRequest.Code)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	iCallback := &inputdata.Callback{
		Request:    iCallbackRequest,
		OAuthToken: oauthToken,
	}
	return ct.inputport.Callback(iCallback)
}
