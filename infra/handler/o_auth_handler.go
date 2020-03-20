package handler

import (
	"cln-arch/interface/controller"
	inputdata "cln-arch/usecase/input/data"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (h *OAuthHandler) Auth(c *Context) error {
	login, err := h.controller.Auth()
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
		return err
	}
	c.Response().Header().Set("Location", login.URL)
	return c.JSON(http.StatusTemporaryRedirect, login)
}

func (h *OAuthHandler) Callback(c *Context) error {
	callback := &inputdata.CallbackRequest{}
	c.Bind(callback)
	info, err := h.controller.Callback(callback)
	if err != nil {
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, info)
}
