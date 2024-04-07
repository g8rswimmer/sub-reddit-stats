package convert

import (
	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
)

// SubredditListingToProto will convert the datastore subreddit listing data to
// the protobuf message.
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
