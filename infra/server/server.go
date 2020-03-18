package server

import (
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"cln-arch/infra/handler"
)

func Start() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(
		middleware.Recover(),
		func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&handler.Context{
					Context: c,
				})
			}
		},
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: logFormat(),
			Output: os.Stdout,
		}),
	)

	// Handlers
	oauthHandler := handler.NewOAuthHandler()

	auth := e.Group("/auth")
	github := auth.Group("/github")
	github.GET("/login", c(oauthHandler.Login))
	github.GET("/callback", c(oauthHandler.Callback))
	github.GET("/token", c(oauthHandler.Auth))

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

func logFormat() string {
	// Refer to https://github.com/tkuchiki/alp
	var format string
	strings.Join([]string{
		"time:${time_rfc3339}\t\n",
		"host:${remote_ip}\t\n",
		"forwardedfor:${header:x-forwarded-for}\t\n",
		"req:-\t\n",
		"status:${status}\t\n",
		"method:${method}\t\n",
		"uri:${uri}\t\n",
		"size:${bytes_out}\t\n",
		"referer:${referer}\t\n",
		"ua:${user_agent}\t\n",
		"reqtime_ns:${latency}\t\n",
		"cache:-\t\n",
		"runtime:-\t\n",
		"apptime:-\t\n",
		"vhost:${host}\t\n",
		"reqtime_human:${latency_human}\t\n",
		"x-request-id:${id}\t\n",
		"host:${host}\n",
	}, "")

	return format
}
