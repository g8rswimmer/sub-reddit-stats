package worker

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/reddit"
)

const (
	limit      = 100
	ctxTO      = 10 * time.Second
	maxRetries = 10
)

var errRetrieveListing = errors.New("error retrieving listing")

type Presister interface {
	Store(ctx context.Context, children []datastore.SubredditListing) error
}

type RedditLister interface {
	SubredditListingNew(ctx context.Context, subreddit string, params ...reddit.Params) (*model.RedditListing, error)
}

type Runner struct {
	Lister    RedditLister
	Presister Presister
	done      chan struct{}
	after     string
}

func (r *Runner) Start(subreddit string) <-chan error {
	to := time.NewTimer(10 * time.Millisecond)
	r.done = make(chan struct{})
	routineErr := make(chan error)

	go func() {
		for {
			select {
			case <-to.C:
				sleep, err := r.process(subreddit)
				if err != nil {
					routineErr <- err
					return
				}
				to.Reset(sleep)
			case <-r.done:
				return
			}
		}
	}()
	return routineErr
}

func (r *Runner) process(subreddit string) (time.Duration, error) {
	listing, err := r.handleSubredditListing(subreddit)
	if err != nil {
		return 0, err
	}
	r.after = listing.Data.After
	if err := r.presistSubredditListings(listing.Data.Children); err != nil {
		return 0, err
	}
	backoff := 1 * time.Second
	if listing.RateLimiting.Remaining > 0 && listing.RateLimiting.Reset > 0 {
		slog.Info("backoff", "remaining", listing.RateLimiting.Remaining, "reset", listing.RateLimiting.Reset)
		backoff = listing.RateLimiting.Reset / time.Duration(listing.RateLimiting.Remaining)
		backoff += time.Millisecond
	}
	return backoff, nil
}

func (r *Runner) presistSubredditListings(children []model.SubredditChild) error {
	ctx, cancel := context.WithTimeout(context.Background(), ctxTO)
	defer cancel()
	listings := make([]datastore.SubredditListing, len(children))
	for i, c := range children {
		listings[i] = datastore.SubredditListing(c.Data)
	}
	return r.Presister.Store(ctx, listings)
}

func (r *Runner) handleSubredditListing(subreddit string) (*model.RedditListing, error) {
	retries := maxRetries
	for retries > 0 {
		var httpErr *reddit.HTTPError

		backoff := time.Millisecond * 100

		listing, err := r.subredditListings(subreddit)
		switch {
		case err == nil:
			return listing, nil
		case errors.As(err, &httpErr):
			if httpErr.StatusCode == http.StatusTooManyRequests {
				backoff = httpErr.RateLimiting.Reset
				if backoff == 0 {
					backoff = 30 * time.Second
				}
			}
		default:

		}
		slog.Error("http error when getting listings", "error", httpErr.Error())
		retries--
		time.Sleep(backoff)
	}
	return nil, errRetrieveListing
}
func (r *Runner) subredditListings(subreddit string) (*model.RedditListing, error) {
	params := []reddit.Params{reddit.WithLimit(limit)}
	if len(r.after) != 0 {
		params = append(params, reddit.WithAfter(r.after))
	}
	ctx, cancel := context.WithTimeout(context.Background(), ctxTO)
	defer cancel()

	return r.Lister.SubredditListingNew(ctx, subreddit, params...)
}

func (r *Runner) Shutdown() {
	close(r.done)
	slog.Info("sutting down the runner")
}
