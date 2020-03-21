package server

import (
	"cln-arch/infra/handler"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
		}),
		middleware.Recover(),
		func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&handler.Context{
					Context: c,
				})
			}
		},
	)

	// Handlers
	oauthHandler := handler.NewOAuthHandler()

	auth := e.Group("/auth")
	github := auth.Group("/github")
	github.GET("/auth", c(oauthHandler.Auth))
	github.GET("/callback", c(oauthHandler.Callback))

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
