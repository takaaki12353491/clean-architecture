package server

import "github.com/labstack/echo/v4"

type (
	Context struct {
		echo.Context
	}

	callFunc func(c *Context) error
)

func c(h callFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(*Context))
	}
}
