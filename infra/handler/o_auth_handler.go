package handler

import (
	"github.com/labstack/echo/v4"

	"cln-arch/interface/controller"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (h *OAuthHandler) Login(c echo.Context) error {
	return nil
}

func (h *OAuthHandler) Callback(c echo.Context) error {
	return nil
}

func (h *OAuthHandler) Auth(c echo.Context) error {
	return nil
}
