package controller

import (
	"cln-arch/config"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputdata "cln-arch/usecase/output/data"
	"context"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
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

func (ct *OAuthController) Callback(context context.Context, iCallbackRequest *inputdata.CallbackRequest) (*outputdata.Callback, error) {
	githubConf := config.NewGithubConf()
	token, err := githubConf.Exchange(context, iCallbackRequest.Code)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	client := github.NewClient(oauth2.NewClient(context, oauth2.StaticTokenSource(token)))
	u, _, err := client.Users.Get(context, "")
	if err != nil {
		panic(err)
	}
	iUser := &inputdata.GithubUser{
		ID:        uint(u.GetID()),
		Name:      u.GetLogin(),
		AvatarURL: u.GetAvatarURL(),
	}
	iCallback := &inputdata.Callback{
		Request:    iCallbackRequest,
		User:       iUser,
		OAuthToken: token,
	}
	return ct.inputport.Callback(iCallback)
}
