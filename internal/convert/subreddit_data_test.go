package convert

import (
	"testing"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/stretchr/testify/assert"
)

func TestSubredditListingToProto(t *testing.T) {
	type args struct {
		srl datastore.SubredditListing
	}
	tests := []struct {
		name string
		args args
		want *redditv1.SubredditData
	}{
		{
			name: "convert",
			args: args{
				srl: datastore.SubredditListing{
					Title:               "Registering my kid for kindergarten...Do you think they'd honor it? 😂",
					Downs:               0,
					UpvoteRatio:         1.0,
					Ups:                 1,
					TotalAwardsReceived: 0,
					Name:                "t3_1bv8ijk",
					Subreddit:           "funny",
					ID:                  "1bv8ijk",
					Author:              "dbzcat",
				},
			},
			want: &redditv1.SubredditData{
				Title:               "Registering my kid for kindergarten...Do you think they'd honor it? 😂",
				Downs:               0,
				UpvoteRatio:         1.0,
				Ups:                 1,
				TotalAwardsReceived: 0,
				Name:                "t3_1bv8ijk",
				Subreddit:           "funny",
				Id:                  "1bv8ijk",
				Author:              "dbzcat",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubredditListingToProto(tt.args.srl)
			assert.Equal(t, tt.want, got)
		})
	}
}
