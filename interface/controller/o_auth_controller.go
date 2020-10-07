package controller

import (
	"cln-arch/config"
	"cln-arch/consts"
	"cln-arch/errs"
	"cln-arch/infra/database"
	"cln-arch/interface/presenter"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	"cln-arch/usecase/interactor"
	"context"
	"net/http"

	github "github.com/google/go-github/github"
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
// @router /oauth [post]
func (ctrl *OAuthController) Auth(c Context) error {
	oAuth, err := ctrl.inputport.Auth()
	service := c.QueryParam(serviceQP)
	oauthConfig, err := config.OAuthConfig(service)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	url := oauthConfig.AuthCodeURL(oAuth.State)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// Callback ...
// @summary
// @description
// @tags OAuth
// @accept json
// @produce json
// @success 200 {object} outputdata.Callback ""
// @failure 400 {string} string ""
// @router /oauth/callback [get]
func (ctrl *OAuthController) Callback(c Context) error {
	request := &CallbackRequest{}
	err := c.Bind(request)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	service := c.QueryParam(serviceQP)
	ctx := c.CTX()
	iCallback, err := ctrl.iCallback(ctx, service, request)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	oCallback, err := ctrl.inputport.Callback(iCallback)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, oCallback)
}

func (ctrl *OAuthController) iCallback(ctx context.Context, service string, request *CallbackRequest) (*inputdata.Callback, error) {
	oauthConfig, err := config.OAuthConfig(service)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token, err := oauthConfig.Exchange(ctx, request.Code)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var iUser *inputdata.User
	switch service {
	case consts.Github:
		client := github.NewClient(oauth2.NewClient(ctx, oauth2.StaticTokenSource(token)))
		u, _, err := client.Users.Get(ctx, "")
		if err != nil {
			log.Error(err)
			return nil, err
		}
		iUser = &inputdata.User{
			ID:        uint(u.GetID()),
			Name:      u.GetLogin(),
			AvatarURL: u.GetAvatarURL(),
		}
	default:
		err = errs.Invalidated.New("unavailable service")
		return nil, err
	}
	iCallback := &inputdata.Callback{
		Code:       request.Code,
		State:      request.State,
		User:       iUser,
		OAuthToken: token,
	}
	return iCallback, nil
}
