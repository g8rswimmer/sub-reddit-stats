package manager

import (
	"context"
	"errors"

	"github.com/g8rswimmer/sub-reddit-stats/internal/convert"
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/errorx"
	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

type Fetcher interface {
	SubredditUps(ctx context.Context, subreddit string, limit int) ([]datastore.SubredditListing, error)
	SubredditPosts(ctx context.Context, subreddit string, limit int) ([]model.SubredditPost, error)
}

type Reddit struct {
	Fetcher Fetcher
}

func (r *Reddit) SubredditMostUps(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditData, error) {

	data, err := r.Fetcher.SubredditUps(ctx, subreddit, limit)
	if err != nil {
		return nil, errors.Join(err, errorx.ErrDatabase)
	}
	ups := make([]*redditv1.SubredditData, len(data))
	for i, d := range data {
		ups[i] = convert.SubredditListingToProto(d)
	}
	return ups, nil
}

func (r *Reddit) SubredditAuthorPosts(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditPost, error) {
	data, err := r.Fetcher.SubredditPosts(ctx, subreddit, limit)
	if err != nil {
		return nil, errors.Join(err, errorx.ErrDatabase)
	}
	posts := make([]*redditv1.SubredditPost, len(data))
	for i, d := range data {
		posts[i] = convert.SubredditPostToProto(d)
	}
	return posts, nil
}
