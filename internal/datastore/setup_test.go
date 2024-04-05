package datastore

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseSetup() (*sqlx.DB, error) {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory")
	if err != nil {
		return nil, fmt.Errorf("unable to open database: %w", err)
	}
	migration := &Migration{
		DB: db,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := migration.Apply(ctx); err != nil {
		return nil, fmt.Errorf("unable to migrate database: %w", err)
	}

	return sqlx.NewDb(db, "sqlite3"), nil
}
