package main

import (
	"context"
	"flag"
	"log/slog"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/config"
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
)

func main() {
	cfgFileName := flag.String("config", "", "config file for migration")
	flag.Parse()

	cfg, err := config.SettingFromFile(*cfgFileName)
	if err != nil {
		slog.Error("unable to load configuration settings", "error", err.Error())
		panic(err)
	}

	db, err := datastore.Open(cfg.Database.DataSource)
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
