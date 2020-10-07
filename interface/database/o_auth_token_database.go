package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type OAuthTokenDatabase struct {
	*gorm.DB
}

func NewOAuthTokenDatabase() repository.OAuthTokenRepository {
	return &OAuthTokenDatabase{NewConnection()}
}

func (db *OAuthTokenDatabase) Store(token *model.OAuthToken) error {
	if err := db.Create(token).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
