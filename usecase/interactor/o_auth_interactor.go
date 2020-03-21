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
	outputport      outputport.OAuthOutputPort
	oauthRepository repository.OAuthRepository
}

func NewOAuthInteractor(
	outputport outputport.OAuthOutputPort,
	oauthRepository repository.OAuthRepository,
) inputport.OAuthInputPort {
	return &OAuthInteractor{
		outputport:      outputport,
		oauthRepository: oauthRepository,
	}
}

// Auth is ...
func (it *OAuthInteractor) Auth() (*outputdata.Auth, error) {
	state, err := model.NewOAuthState()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.oauthRepository.StoreState(state)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Auth(state), nil
}

func (it *OAuthInteractor) Callback(iCallback *inputdata.Callback) (*outputdata.Callback, error) {
	state, err := it.oauthRepository.FindStateByState(iCallback.Request.State)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if state.Expiry.After(time.Now()) {
		errMsg := "state is expiry"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	user, err := model.NewUser()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	token, err := model.NewOAuthToken(user, iCallback.OAuthToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = it.oauthRepository.StoreToken(token)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Callback(token), nil
}
