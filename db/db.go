package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// todo receive env data
func NewDatabase(user string, dbname string, password string) (*sql.DB, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
