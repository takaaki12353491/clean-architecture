package interactor

import (
	"cln-arch/domain/model"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputdata "cln-arch/usecase/output/data"
	outputport "cln-arch/usecase/output/port"
	"cln-arch/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// OAuthInteractor is ...
type OAuthInteractor struct {
	outputport      outputport.OAuthOutputPort
	userRepository  repository.UserRepository
	stateRepository repository.OAuthStateRepository
	tokenRepository repository.OAuthTokenRepository
}

func NewOAuthInteractor(
	outputport outputport.OAuthOutputPort,
	userRepository repository.UserRepository,
	stateRepository repository.OAuthStateRepository,
	tokenRepository repository.OAuthTokenRepository,
) inputport.OAuthInputPort {
	return &OAuthInteractor{
		outputport:      outputport,
		userRepository:  userRepository,
		stateRepository: stateRepository,
		tokenRepository: tokenRepository,
	}
}

// Auth is ...
func (it *OAuthInteractor) Auth() (*outputdata.Auth, error) {
	state, err := model.NewOAuthState()
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	err = it.stateRepository.Store(state)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return it.outputport.Auth(state), nil
}

func (it *OAuthInteractor) Callback(iCallback *inputdata.Callback) (*outputdata.Callback, error) {
	state, err := it.stateRepository.FindByState(iCallback.Request.State)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	err = it.stateRepository.Delete(state)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	existingUser, _ := it.userRepository.FindByID(iCallback.User.ID)
	if existingUser != nil {
		log.WithFields(log.Fields{}).Infoln("The user already exists")
		return it.outputport.Callback(existingUser), nil
	}
	user, err := model.NewUser(iCallback.User.ID, iCallback.User.Name)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	user.AvatorURL = iCallback.User.AvatarURL
	err = it.userRepository.Store(user)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	token, err := model.NewOAuthToken(user, iCallback.OAuthToken)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	err = it.tokenRepository.Store(token)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return it.outputport.Callback(user), nil
}
