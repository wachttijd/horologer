package database

import (
	"database/sql"
	"horologer/conf"
	"horologer/models"
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

func AddNewStrongbox(box models.Strongbox) error {
	_, err := Db.Exec(`
		INSERT INTO strongboxes (
			general_id,
			available_after,
			decryption_key,
			data
		) VALUES (?, ?, ?, ?)
	`,
		box.GeneralId,
		box.AvailableAfter,
		box.DecryptionKey,
		box.Data,
	)

	return err
}

func GetStrongbox(generalId string) (models.Strongbox, error) {
	var box models.Strongbox

	err := Db.QueryRow(`
		SELECT 
			general_id,
			available_after,
			decryption_key,
			data
		FROM
			strongboxes
		WHERE
			general_id = ?
	`, generalId).Scan(
		&box.GeneralId,
		&box.AvailableAfter,
		&box.DecryptionKey,
		&box.Data,
	)

	return box, err
}
