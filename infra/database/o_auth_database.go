package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type OAuthDatabase struct {
	*SQLHandler
}

func NewOAuthDatabase() repository.OAuthRepository {
	return &OAuthDatabase{SQLHandler: NewSQLHandler()}
}

func (db *OAuthDatabase) FindStateByState(state string) (*model.OAuthState, error) {
	oauthState := &model.OAuthState{}
	if err := db.First(oauthState, "state = ?", state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return oauthState, nil
}

func (db *OAuthDatabase) FindTokenByToken(token string) (*model.OAuthToken, error) {
	oauthToken := &model.OAuthToken{}
	if err := db.First(oauthToken, "access_token = ?", token).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return oauthToken, nil
}

func (db *OAuthDatabase) StoreState(state *model.OAuthState) error {
	if err := db.Create(state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

func (db *OAuthDatabase) StoreToken(token *model.OAuthToken) error {
	if err := db.Create(token).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
