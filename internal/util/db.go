package util

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
)

func CreateTables(db *sql.DB) error {
	log.Info("[db] Creating tables")
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		password TEXT,
		email TEXT		
	)`)
	return err
}
