package datastore

import (
	"context"
	"fmt"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/jmoiron/sqlx"
)

const defaultLimit = 5

type Listing struct {
	DB *sqlx.DB
}

func (p *Listing) Store(ctx context.Context, children []model.SubredditChild) error {

	listings := make([]model.SubredditData, len(children))
	for i := range children {
		listings[i] = children[i].Data
	}

	if _, err := p.DB.NamedExecContext(ctx, insertListing, listings); err != nil {
		return fmt.Errorf("insert of listings %d to table: %w", len(listings), err)
	}
	return nil
}

func (p *Listing) SubredditUps(ctx context.Context, subreddit string, limit int) ([]model.SubredditData, error) {
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

	listings := []model.SubredditData{}
	for rows.Next() {
		data := model.SubredditData{}
		if err := rows.StructScan(&data); err != nil {
			return nil, fmt.Errorf("listing ups row scan error: %w", err)
		}
		listings = append(listings, data)
	}

	return listings, nil
}
