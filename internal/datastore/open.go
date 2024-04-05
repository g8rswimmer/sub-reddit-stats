package datastore

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Open(dataSourceName string) (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", dataSourceName)
}
