package datastore

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

const defaultLimit = 5

// Listing will handle datastore interface for the subreddit listing
// data.
type Listing struct {
	DB *sqlx.DB
}

// Store will store the data listing from a subreddit
func (p *Listing) Store(ctx context.Context, listings []SubredditListing) error {

	if _, err := p.DB.NamedExecContext(ctx, insertListing, listings); err != nil {
		return fmt.Errorf("insert of listings %d to table: %w", len(listings), err)
	}
	return nil
}

// SubredditUps will return the subreddit posts with the most ups (up votes).
// The limit is optional and if none are provided, a default will be used.
func (p *Listing) SubredditUps(ctx context.Context, subreddit string, limit int) ([]SubredditListing, error) {
	if limit <= 0 {
		limit = defaultLimit
	}
	rows, err := p.DB.NamedQueryContext(ctx, listingUps, map[string]any{
		"subreddit": subreddit,
		"limit":     limit,
	})
	if err != nil {
		return nil, fmt.Errorf("listing ups for %s with limit %d: %w", subreddit, limit, err)
	}
	defer rows.Close()

	listings := []SubredditListing{}
	for rows.Next() {
		data := SubredditListing{}
		if err := rows.StructScan(&data); err != nil {
			return nil, fmt.Errorf("listing ups row scan error: %w", err)
		}
		listings = append(listings, data)
	}

	return listings, nil
}

// SubredditPosts will return the authors with the most posts for a given subreddit.
// The limit is optional and if none are provided, a default will be used.
func (p *Listing) SubredditPosts(ctx context.Context, subreddit string, limit int) ([]SubredditPost, error) {
	if limit <= 0 {
		limit = defaultLimit
	}
	rows, err := p.DB.NamedQueryContext(ctx, listingAuthorPosts, map[string]any{
		"subreddit": subreddit,
		"limit":     limit,
	})
	if err != nil {
		return nil, fmt.Errorf("listing posts for %s with limit %d: %w", subreddit, limit, err)
	}
	defer rows.Close()

	posts := []SubredditPost{}
	for rows.Next() {
		data := SubredditPost{}
		if err := rows.StructScan(&data); err != nil {
			return nil, fmt.Errorf("listing posts row scan error: %w", err)
		}
		posts = append(posts, data)
	}
	return posts, nil
}
