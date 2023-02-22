package database

import (
	"database/sql"
	"horologer/conf"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	DSN, err := conf.DatabaseDSN()

	if err != nil {
		log.Fatal(err)
	}

	var DbInitErr error

	Db, DbInitErr = sql.Open("mysql", DSN)

	if DbInitErr != nil {
		log.Fatal(DbInitErr)
	}

	if err := Db.Ping(); err != nil {
		log.Fatal(err)
	}
}
