package server

import (
	"cln-arch/infra/handler"

	"github.com/labstack/echo/v4"
)

type callFunc func(c *handler.Context) error

func c(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*handler.Context))
	}
}
