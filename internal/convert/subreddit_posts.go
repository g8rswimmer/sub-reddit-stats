package convert

import (
	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

func SubredditPostToProto(m model.SubredditPost) *redditv1.SubredditPost {
	return &redditv1.SubredditPost{
		AuthorFullname: m.AuthorFullname,
		Author:         m.Author,
		Posts:          int32(m.Posts),
	}
}
