package manager

import (
	"context"
	"errors"

	"github.com/g8rswimmer/sub-reddit-stats/internal/convert"
	"github.com/g8rswimmer/sub-reddit-stats/internal/errorx"
	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

type Fetcher interface {
	ListingUps(ctx context.Context, subreddit string, limit int) ([]model.SubredditData, error)
}

type Reddit struct {
	Fetcher Fetcher
}

func (r *Reddit) SubredditMostUps(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditData, error) {

	data, err := r.Fetcher.ListingUps(ctx, subreddit, limit)
	if err != nil {
		return nil, errors.Join(err, errorx.ErrDatabase)
	}
	ups := make([]*redditv1.SubredditData, len(data))
	for i, d := range data {
		ups[i] = convert.SubredditDataToProto(d)
	}
	return ups, nil
}
