package convert

import (
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

func SubredditPostToProto(srp datastore.SubredditPost) *redditv1.SubredditPost {
	return &redditv1.SubredditPost{
		AuthorFullname: srp.AuthorFullname,
		Author:         srp.Author,
		Posts:          int32(srp.Posts),
	}
}
