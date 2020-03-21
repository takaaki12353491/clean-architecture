package database

import (
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type SQLHandler struct {
	*gorm.DB
}

func NewSQLHandler() *SQLHandler {
	DBMS := "mysql"
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PROTOCOL := "tcp"
	DBNAME := "cln_arch"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Panicln(err)
	}
	return &SQLHandler{DB: db}
}
