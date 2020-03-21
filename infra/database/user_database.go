package database

import (
	"cln-arch/domain/model"
	"cln-arch/usecase/repository"

	log "github.com/sirupsen/logrus"
)

type UserDatabase struct {
	*SQLHandler
}

func NewUserDatabase() repository.UserRepository {
	return &UserDatabase{SQLHandler: NewSQLHandler()}
}

func (db *UserDatabase) FindByID(id uint) (*model.User, error) {
	user := new(model.User)
	user.ID = id
	if err := db.First(user).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return user, nil
}

func (db *UserDatabase) Store(user *model.User) error {
	if err := db.Create(user).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

func (db *UserDatabase) Update(user *model.User) error {
	if err := db.Save(user).Error; err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
