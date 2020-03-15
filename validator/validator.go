package validator

import (
	"cln-arch/errs"
	"net/http"

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

func Validate(i interface{}) errs.HTTPError {
	err := NewValidator().validator.Struct(i)
	if err != nil {
		return errs.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
