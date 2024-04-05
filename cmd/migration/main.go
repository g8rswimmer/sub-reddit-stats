package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
)

func main() {
	db, err := datastore.Open("./db/sqlite-database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	migration := &datastore.Migration{
		DB: db.DB,
	}
	slog.Info("applying migrations")
	if err := migration.Apply(ctx); err != nil {
		panic(err)
	}
	slog.Info("done applying migrations")
}
