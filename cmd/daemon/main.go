package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/config"
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/oauth"
	"github.com/g8rswimmer/sub-reddit-stats/internal/reddit"
	"github.com/g8rswimmer/sub-reddit-stats/internal/worker"
)

func main() {
	cfgFileName := flag.String("config", "", "config file for migration")
	flag.Parse()

	cfg, err := config.SettingFromFile(*cfgFileName)
	if err != nil {
		slog.Error("unable to load configuration settings", "error", err.Error())
		panic(err)
	}

	slog.Info("starting daemon init....")
	db, err := datastore.Open(cfg.Database.DataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	manager, err := oauth.NewManager(context.Background(), oauth.WithCredentials(cfg.Reddit.ClientID, cfg.Reddit.ClientSecret), oauth.WithBaseURL(cfg.Reddit.OAuthURL))
	if err != nil {
		slog.Error("unable to run oauth manager", "error", err.Error())
		panic(err)
	}
	defer manager.Shutdown()

	redditClient := &reddit.Client{
		BaseURL: cfg.Reddit.BaseURL,
		Auth:    manager,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	presist := &datastore.Listing{
		DB: db,
	}

	runner := worker.Runner{
		Lister:    redditClient,
		Presister: presist,
	}
	defer runner.Shutdown()

	slog.Info("starting daemon runner...")
	runnerErr := runner.Start(cfg.Reddit.Subreddit)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case e := <-runnerErr:
		slog.Error("runner error", "error", e.Error())
	case <-sigs:
	}
	slog.Info("daemon stopped....")

}
