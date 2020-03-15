package handler

import (
	"cln-arch/infra/server"
	"cln-arch/interface/controller"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (h *OAuthHandler) Login(c server.Context) error {
	return nil
}

func (h *OAuthHandler) Callback(c server.Context) error {
	return nil
}

func (h *OAuthHandler) Auth(c server.Context) error {
	return nil
}
