package model

import (
	"cln-arch/validator"

	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
)

type User struct {
	gorm.Model
	Name      string `validate:"required"`
	AvatorURL string
}

func NewUser(id uint, name string) (*User, error) {
	user := &User{
		Model: gorm.Model{ID: id},
		Name:  name,
	}
	err := validator.Validate(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}
