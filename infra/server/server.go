package server

import (
	"cln-arch/interface/controller"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	// Echo instance
	e := NewEcho()
	// Middleware
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
		}),
		middleware.Recover(),
		func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&Context{
					Context: c,
				})
			}
		},
	)

	// Controllers
	oauthController := controller.NewOAuthController()

	auth := e.EchoGroup("/auth")
	github := auth.EchoGroup("/github")
	github.GET("/auth", oauthController.Auth)
	github.GET("/callback", oauthController.Callback)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func logFormat() string {
	format := strings.Replace(middleware.DefaultLoggerConfig.Format, ",", ",\n  ", -1)
	format = strings.Replace(format, "{", "{\n  ", 1)
	format = strings.Replace(format, "}}", "}\n}", 1)
	return format
}
