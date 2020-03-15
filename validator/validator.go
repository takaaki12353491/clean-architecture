package validator

import (
	"cln-arch/errs"

	"gopkg.in/go-playground/validator.v9"
)

type Validator struct {
	validator *validator.Validate
}

func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

func Validate(i interface{}) error {
	err := NewValidator().validator.Struct(i)
	if err != nil {
		return errs.Invalidated.Wrap(err, err.Error())
	}
	return nil
}
