package convert

import (
	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

func SubredditDataToProto(m model.SubredditData) *redditv1.SubredditData {
	return &redditv1.SubredditData{
		Id:                  m.ID,
		Title:               m.Title,
		Ups:                 int32(m.Ups),
		Downs:               int32(m.Downs),
		UpvoteRatio:         float32(m.UpvoteRatio),
		TotalAwardsReceived: int32(m.TotalAwardsReceived),
		Name:                m.Name,
		Subreddit:           m.Subreddit,
		Author:              m.Author,
		AuthorFullname:      m.AuthorFullname,
	}
}
