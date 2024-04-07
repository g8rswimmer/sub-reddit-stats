package convert

import (
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
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

func SubredditListingToProto(srl datastore.SubredditListing) *redditv1.SubredditData {
	return &redditv1.SubredditData{
		Id:                  srl.ID,
		Title:               srl.Title,
		Ups:                 int32(srl.Ups),
		Downs:               int32(srl.Downs),
		UpvoteRatio:         float32(srl.UpvoteRatio),
		TotalAwardsReceived: int32(srl.TotalAwardsReceived),
		Name:                srl.Name,
		Subreddit:           srl.Subreddit,
		Author:              srl.Author,
		AuthorFullname:      srl.AuthorFullname,
	}
}
