package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// todo receive env data
func NewDatabase(user string, dbname string, password string) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
