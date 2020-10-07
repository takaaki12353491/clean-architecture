package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type OAuthStateDatabase struct {
	*gorm.DB
}

func NewOAuthStateDatabase() repository.OAuthStateRepository {
	return &OAuthStateDatabase{NewConnection()}
}

func (db *OAuthStateDatabase) FindByState(str string) (*model.OAuthState, error) {
	state := &model.OAuthState{}
	if err := db.First(state, "state = ?", str).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return state, nil
}

func (db *OAuthStateDatabase) Store(state *model.OAuthState) error {
	if err := db.Create(state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

func (db *OAuthStateDatabase) Delete(state *model.OAuthState) error {
	if err := db.Unscoped().Delete(state).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
