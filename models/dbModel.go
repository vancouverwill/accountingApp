package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type dbModels struct {
}

var myDb = dbModels{}

const (
	address  string = "localhost:3306"
	password string = ""
	username string = "root"
)

func (myDb dbModels) setup() (*sql.DB, error) {
	db, e := sql.Open("mysql", username+":"+password+"@tcp("+address+")/accountancyApp")
	return db, e
}
