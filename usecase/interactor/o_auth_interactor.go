package interactor

import (
	log "github.com/sirupsen/logrus"

	"cln-arch/errs"
	inputdata "cln-arch/usecase/input/data"
	inputport "cln-arch/usecase/input/port"
	outputdata "cln-arch/usecase/output/data"
	outputport "cln-arch/usecase/output/port"
	"cln-arch/usecase/repository"
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
func (it *OAuthInteractor) Login(oauth *inputdata.Login) (*outputdata.Login, error) {
	err := it.oauthRepository.StoreState(oauth.State, oauth.Session.ID, oauth.Expiry)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return it.outputport.Login(oauth.State, oauth.URL), nil
}

func (it *OAuthInteractor) Callback(callback *inputdata.Callback) (*outputdata.Callback, error) {
	// recieved state is expected or not
	state, err := it.oauthRepository.FindBySessionID(callback.ID)
	if err != nil {
		return nil, err
	}
	if state != callback.State {
		return nil, errs.Forbidden.New("not match state")
	}
	id, err := it.oauthRepository.StoreGithubToken(callback.GithubToken)
	if err != nil {
		return nil, err
	}
	count, err := it.oauthRepository.StoreUserToken(callback.ID, callback.UserToken, id)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errs.Forbidden.New("no user session")
	}
	return it.outputport.Callback(callback.UserToken.Token), nil
}
