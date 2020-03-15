package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"cln-arch/infra/handler"
)

func Start() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&Context{c})
			}
		},
	)

	oauthHandler := handler.NewOAuthHandler()

	auth := e.Group("/auth")
	github := auth.Group("/github")
	github.GET("/login", oauthHandler.Login)
	github.GET("/callback", oauthHandler.Callback)
	github.GET("/token", oauthHandler.Auth)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}