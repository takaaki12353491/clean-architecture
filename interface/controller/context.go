package controller

import "context"

type Context interface {
	CTX() context.Context
	String(code int, s string) error
	JSON(code int, i interface{}) error
	Redirect(code int, url string) error
	Bind(i interface{}) error
}
