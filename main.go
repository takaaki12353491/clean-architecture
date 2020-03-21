package main

import (
	_ "cln-arch/docs"
	"cln-arch/infra/server"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

// @title CLN-ARCH API
// @version 1.0
// @description This is a personal project
// @license.name takaaki12353491
// @host localhost:8080
// @BasePath /
func main() {
	// Log settings
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)

	server.Start()
}
