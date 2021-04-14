package database

import (
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/mattn/go-sqlite3"
)

const DatabaseName = "./software.db"

func GetConnection() *sqlx.DB {
	conn, err := sqlx.Connect("sqlite3", DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database connection created")
	return conn
}
