package controller

import (
	"cln-arch/config"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputdata "cln-arch/usecase/output/data"
	"context"

	log "github.com/sirupsen/logrus"
)

type OAuthController struct {
	inputport inputport.OAuthInputPort
}

func NewOAuthController(inputport inputport.OAuthInputPort) *OAuthController {
	return &OAuthController{inputport: inputport}
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
