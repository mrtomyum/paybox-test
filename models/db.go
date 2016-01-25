package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func InitDB(conn string) {
	var err error
	db, err = sql.Open("sqlite3", conn)

	if err != nil {
		log.Panic("Cannot Open DB: ", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic("Ping error:", err)
	}
}
