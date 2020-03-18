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

	"github.com/google/uuid"
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

func (ct *OAuthController) Login(iSession *inputdata.Session) (*outputdata.Login, error) {
	userID := uuid.New().String()
	state := createRand()
	serverConf := config.NewServer()
	url := serverConf.Github.AuthCodeURL(state)
	expiry := time.Now().Add(10 * time.Minute)
	iLogin := &inputdata.Login{
		UserID:  userID,
		State:   state,
		URL:     url,
		Session: iSession,
		Expiry:  &expiry,
	}
	return ct.inputport.Login(iLogin)
}

func (ct *OAuthController) Callback(iCallbackRequest *inputdata.CallbackRequest) (*outputdata.Callback, error) {
	serverConf := config.NewServer()
	oauthToken, err := serverConf.Github.Exchange(context.Background(), iCallbackRequest.Code)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token := createRand()
	expiry := time.Now().Add(7 * 24 * time.Hour)
	iCallback := &inputdata.Callback{
		Request:    iCallbackRequest,
		OAuthToken: oauthToken,
		Token:      token,
		Expiry:     &expiry,
	}
	return ct.inputport.Callback(iCallback)
}

func (ct *OAuthController) Auth(iAuth *inputdata.Auth) (*outputdata.Auth, error) {
	return ct.inputport.Auth(iAuth)
}
