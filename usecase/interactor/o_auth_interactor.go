package interactor

import (
	"cln-arch/domain/model"
	"cln-arch/errs"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputdata "cln-arch/usecase/output/data"
	outputport "cln-arch/usecase/output/port"
	"cln-arch/usecase/repository"
	"time"

	log "github.com/sirupsen/logrus"
)

// OAuthInteractor is ...
type OAuthInteractor struct {
	oauthRepository repository.OAuthRepository
	outputport      outputport.OAuthOutputPort
}

func NewOAuthInteractor(
	oauthRepository repository.OAuthRepository,
	outputport outputport.OAuthOutputPort,
) inputport.OAuthInputPort {
	return &OAuthInteractor{
		oauthRepository: oauthRepository,
		outputport:      outputport,
	}
}

// Login is ...
func (it *OAuthInteractor) Login(iLogin *inputdata.Login) (*outputdata.Login, error) {
	userState, err := model.NewUserState(iLogin.UserID, iLogin.Session.ID, iLogin.State, iLogin.Expiry)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	login, err := model.NewLogin(iLogin.State, iLogin.URL)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.oauthRepository.StoreUserState(userState)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Login(login), nil
}

func (it *OAuthInteractor) Callback(iCallback *inputdata.Callback) (*outputdata.Callback, error) {
	// recieved state is expected or not
	userState, err := it.oauthRepository.FindUserStateBySessionID(iCallback.Request.Session.ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if userState.State != iCallback.Request.State {
		errMsg := "not match state"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	oauthToken, err := model.NewOAuthToken(userState.UserID, iCallback.OAuthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.oauthRepository.StoreOAuthToken(oauthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	userToken, err := model.NewUserToken(userState.UserID, iCallback.Token, iCallback.Expiry)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.oauthRepository.StoreUserToken(userToken)
	if err != nil {
		return nil, err
	}
	return it.outputport.Callback(userToken), nil
}

func (it *OAuthInteractor) Auth(iAuth *inputdata.Auth) (*outputdata.Auth, error) {
	userState, err := it.oauthRepository.FindUserStateBySessionIDAndUserToken(iAuth.Session.ID, iAuth.Token)
	if userState.Expiry.After(time.Now()) {
		errMsg := "user token expiry"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	oauthToken, err := it.oauthRepository.FindOAuthTokenByUserID(userState.UserID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Auth(oauthToken), nil
}
