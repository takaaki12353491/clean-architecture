package model

import (
	"cln-arch/validator"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type User struct {
	*gorm.Model
}

func NewUser() (*User, error) {
	user := new(User)
	user.ID = uint(uuid.New().ID())
	err := validator.Validate(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
