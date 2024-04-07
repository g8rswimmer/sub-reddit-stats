package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// Open will open the database using the sqlite driver.
func Open(dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", dataSourceName)
}
