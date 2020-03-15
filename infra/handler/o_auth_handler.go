package handler

import (
	"cln-arch/interface/controller"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (h *OAuthHandler) Login(c *Context) error {
	return nil
}

func (h *OAuthHandler) Callback(c *Context) error {
	return nil
}

func (h *OAuthHandler) Auth(c *Context) error {
	return nil
}
