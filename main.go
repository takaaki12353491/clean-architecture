package main

import (
	"cln-arch/infra/router"

	log "github.com/sirupsen/logrus"
)

func main() {
	// Log settings
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)

	router.Start()
}
