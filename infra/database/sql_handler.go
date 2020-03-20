package database

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type SQLHandler struct {
	*gorm.DB
}

func NewSQLHandler() *SQLHandler {
	DBMS := "mysql"
	USER := "root"
	PASS := "####"
	PROTOCOL := "tcp(##.###.##.###:3306)"
	DBNAME := "##"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		log.Panicln(err)
	}
	return &SQLHandler{DB: db}
}
