package controller

import (
	"cln-arch/config"
	"cln-arch/infra/database"
	"cln-arch/interface/presenter"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	"cln-arch/usecase/interactor"
	"net/http"

	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

// CallbackRequest is callback param after github login
type CallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type OAuthController struct {
	inputport inputport.OAuthInputPort
}

func NewOAuthController() *OAuthController {
	return &OAuthController{
		inputport: interactor.NewOAuthInteractor(
			presenter.NewOAuthPresenter(),
			database.NewUserDatabase(),
			database.NewOAuthStateDatabase(),
			database.NewOAuthTokenDatabase(),
		),
	}
}

// Auth ...
// @summary
// @description
// @tags OAuth
// @accept json
// @produce json
// @success 307 {object} outputdata.Auth ""
// @failure 400 {string} string ""
// @router /auth/github/auth [get]
func (ctrl *OAuthController) Auth(c Context) error {
	oAuth, err := ctrl.inputport.Auth()
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.Redirect(http.StatusTemporaryRedirect, oAuth.URL)
}

// Callback ...
// @summary
// @description
// @tags OAuth
// @accept json
// @produce json
// @success 200 {object} outputdata.Callback ""
// @failure 400 {string} string ""
// @router /auth/github/callback [get]
func (ctrl *OAuthController) Callback(c Context) error {
	request := &CallbackRequest{}
	c.Bind(request)
	githubConf := config.NewGithubConf()
	ctx := c.CTX()
	token, err := githubConf.Exchange(ctx, request.Code)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)))
	u, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	iUser := &inputdata.GithubUser{
		ID:        uint(u.GetID()),
		Name:      u.GetLogin(),
		AvatarURL: u.GetAvatarURL(),
	}
	iCallback := &inputdata.Callback{
		Code:       request.Code,
		State:      request.State,
		User:       iUser,
		OAuthToken: token,
	}
	oCallback, err := ctrl.inputport.Callback(iCallback)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, oCallback)
}
