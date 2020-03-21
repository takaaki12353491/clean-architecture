package model

import (
	"cln-arch/validator"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
)

type User struct {
	ID string `gorm:"primary_key"`
}

func NewUser() (*User, error) {
	user := &User{
		ID: uuid.New().String(),
	}
	err := validator.Validate(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
