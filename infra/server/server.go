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
	// Refer to https://github.com/tkuchiki/alp
	var format string
	strings.Join([]string{
		"time:${time_rfc3339},\n",
		"host:${remote_ip},\n",
		"forwardedfor:${header:x-forwarded-for},\n",
		"req:-,\n",
		"status:${status},\n",
		"method:${method},\n",
		"uri:${uri},\n",
		"size:${bytes_out},\n",
		"referer:${referer},\n",
		"ua:${user_agent},\n",
		"reqtime_ns:${latency},\n",
		"cache:-,\n",
		"runtime:-,\n",
		"apptime:-,\n",
		"vhost:${host},\n",
		"reqtime_human:${latency_human},\n",
		"x-request-id:${id},\n",
		"host:${host}\n",
	}, "")

	return format
}
