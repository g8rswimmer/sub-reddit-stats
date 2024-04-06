package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/oauth"
	"github.com/g8rswimmer/sub-reddit-stats/internal/reddit"
	"github.com/g8rswimmer/sub-reddit-stats/internal/worker"
)

func main() {
	slog.Info("starting daemon init....")
	db, err := datastore.Open("./db/sqlite-database.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	manager, err := oauth.NewManager(context.Background(), oauth.WithCredentials("OvUtWulgJ-HGglyUrENANg", "LddtG4rFUULmfEQzdkn83p2gvZw7Aw"))
	if err != nil {
		slog.Error("unable to run oauth manager", "error", err.Error())
		panic(err)
	}
	defer manager.Shutdown()

	redditClient := &reddit.Client{
		BaseURL: "https://oauth.reddit.com",
		Auth:    manager,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
	presist := &datastore.Presister{
		DB: db,
	}

	runner := worker.Runner{
		Lister:    redditClient,
		Presister: presist,
	}
	defer runner.Shutdown()

	slog.Info("starting daemon runner...")
	runnerErr := runner.Start("funny")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case e := <-runnerErr:
		slog.Error("runner error", "error", e.Error())
	case <-sigs:
	}
	slog.Info("daemon stopped....")

}
