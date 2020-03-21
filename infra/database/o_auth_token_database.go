package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type OAuthTokenDatabase struct {
	sql *SQLHandler
}

func NewOAuthTokenDatabase() repository.OAuthTokenRepository {
	return &OAuthTokenDatabase{sql: NewSQLHandler()}
}

func (db *OAuthTokenDatabase) Store(token *model.OAuthToken) error {
	if err := db.sql.Create(token).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
