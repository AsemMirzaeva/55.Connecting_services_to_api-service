package storage

import (
	"database/sql"
	"log"
)

func OpenDb(drname, url string) (*sql.DB, error) {
	db, err := sql.Open(drname, url)
	if err != nil {
		log.Println("Error opening db:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("Error connecting to db:", err)
		return nil, err
	}

	return db, nil
}
