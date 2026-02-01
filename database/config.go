package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// open db conn
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// test db conn
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// set conn pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connection established")

	return db, nil
}
