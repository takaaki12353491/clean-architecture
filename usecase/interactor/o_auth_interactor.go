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

// SetupGithubLogin is ...
func (it *OAuthInteractor) SetupGithubLogin(oauth *inputdata.OAuth) (*outputdata.Login, errs.HTTPError) {
	err := it.oauthRepository.StoreState(oauth.State, oauth.Session.ID, oauth.Expiry)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return it.outputport.Login(oauth.State, oauth.URL), nil
}
