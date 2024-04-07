package manager

import (
	"context"
	"errors"

	"github.com/g8rswimmer/sub-reddit-stats/internal/convert"
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/errorx"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

// Fetcher is used to fetch data from the datastore
type Fetcher interface {
	SubredditUps(ctx context.Context, subreddit string, limit int) ([]datastore.SubredditListing, error)
	SubredditPosts(ctx context.Context, subreddit string, limit int) ([]datastore.SubredditPost, error)
}

// Reddit will manage service calls for subreddit stats
type Reddit struct {
	Fetcher Fetcher
}

// SubredditMostUps will return a list of subreddit listings that have the most ups (up votes).
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

// SubredditAuthorPosts will return a list of authors with the most posts for a given subreddit
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
