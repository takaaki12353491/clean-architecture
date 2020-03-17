package handler

import (
	"cln-arch/interface/controller"
	inputdata "cln-arch/usecase/input/data"
	"net/http"

	"github.com/labstack/gommon/log"
)

type OAuthHandler struct {
	controller controller.OAuthController
}

func NewOAuthHandler() *OAuthHandler {
	return &OAuthHandler{}
}

func (h *OAuthHandler) Login(c *Context) error {
	session := &inputdata.Session{}
	c.Bind(session)
	login, err := h.controller.Login(session)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), err.Error())
	}
	c.Response().Header().Set("Location", login.URL)
	return c.JSON(http.StatusTemporaryRedirect, login)
}

func (h *OAuthHandler) Callback(c *Context) error {
	callback := &inputdata.Callback{}
	c.Bind(callback)
	info, err := h.controller.Callback(callback)
	if err != nil {
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, info)
}

func (h *OAuthHandler) Auth(c *Context) error {
	auth := &inputdata.Auth{}
	c.Bind(auth)
	token, err := h.controller.Auth(auth)
	if err != nil {
		c.String(statusCode(err), err.Error())
		return err
	}
	return c.JSON(http.StatusOK, token)
}
