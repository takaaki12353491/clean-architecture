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
func (it *OAuthInteractor) Login(login *inputdata.Login, _session *inputdata.Session, expiry *time.Time) (*outputdata.Login, error) {
	session := &model.Session{ID: _session.ID}
	err := it.oauthRepository.StoreState(login.State, session, expiry)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return it.outputport.Login(login.State, login.URL), nil
}

func (it *OAuthInteractor) Callback(callback *inputdata.Callback, _githubToken *inputdata.GithubToken, _userToken *inputdata.UserToken) (*outputdata.Callback, error) {
	session := &model.Session{ID: callback.Session.ID}
	githubToken := &model.GithubToken{Token: _githubToken.Token}
	userToken := &model.UserToken{Token: _userToken.Token, Expiry: _userToken.Expiry}
	// recieved state is expected or not
	state, err := it.oauthRepository.FindStateBySession(session)
	if err != nil {
		return nil, err
	}
	if state != callback.State {
		return nil, errs.Forbidden.New("not match state")
	}
	id, err := it.oauthRepository.StoreGithubToken(githubToken)
	if err != nil {
		return nil, err
	}
	count, err := it.oauthRepository.StoreUserToken(session, userToken, id)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errs.Forbidden.New("no user session")
	}
	return it.outputport.Callback(userToken), nil
}

func (it *OAuthInteractor) Auth(auth *inputdata.Auth) (*outputdata.Auth, error) {
	expiry, id, err := it.oauthRepository.FindBySessionIDAndUserToken(auth.Session.ID, auth.Token)
	if expiry.After(time.Now()) {
		return nil, errs.Forbidden.New("user token expiry")
	}
	githubToken, err := it.oauthRepository.FindByUserTokenID(id)
	if err != nil {
		return nil, err
	}
	return it.outputport.Auth(githubToken), nil
}
