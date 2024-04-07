package service

import (
	"context"

	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Manager is used to handle the business logic of the service
type Manager interface {
	SubredditMostUps(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditData, error)
	SubredditAuthorPosts(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditPost, error)
}

// Reddit handles the gRPC service methods
type Reddit struct {
	Manager Manager
}

// GetSubredditMostUps will return the subreddit posts with the most ups (up votes)
// The subreddit parameter is required.
func (r *Reddit) GetSubredditMostUps(ctx context.Context, req *redditv1.GetSubredditMostUpsRequest) (*redditv1.GetSubredditMostUpsResponse, error) {
	if len(req.GetSubreddit()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "subreddit is required")
	}

	resp, err := r.Manager.SubredditMostUps(ctx, req.Subreddit, int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "most ups database error %s", err.Error())
	}

	return &redditv1.GetSubredditMostUpsResponse{
		SubredditPosts: resp,
	}, nil
}

// GetSubredditAuthorPosts will return author's with most posts for a given subreddit
// The subreddit parameter is required.
func (r *Reddit) GetSubredditAuthorPosts(ctx context.Context, req *redditv1.GetSubredditAuthorPostsRequest) (*redditv1.GetSubredditAuthorPostsResponse, error) {
	if len(req.GetSubreddit()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "subreddit is required")
	}
	resp, err := r.Manager.SubredditAuthorPosts(ctx, req.Subreddit, int(req.Limit))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "most ups database error %s", err.Error())
	}

	return &redditv1.GetSubredditAuthorPostsResponse{
		AuthorPosts: resp,
	}, nil
}
