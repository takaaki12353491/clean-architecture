package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type OAuthStateDatabase struct {
	sql *SQLHandler
}

func NewOAuthStateDatabase() repository.OAuthStateRepository {
	return &OAuthStateDatabase{sql: NewSQLHandler()}
}

func (db *OAuthStateDatabase) FindByState(str string) (*model.OAuthState, error) {
	state := &model.OAuthState{}
	if err := db.sql.First(state, "state = ?", state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return state, nil
}

func (db *OAuthStateDatabase) Store(state *model.OAuthState) error {
	if err := db.sql.Create(state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

func (db *OAuthStateDatabase) Delete(state *model.OAuthState) error {
	if err := db.sql.Delete(state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
