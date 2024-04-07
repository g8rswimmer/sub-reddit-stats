package convert

import (
	"testing"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/stretchr/testify/assert"
)

func TestSubredditPostToProto(t *testing.T) {
	type args struct {
		srp datastore.SubredditPost
	}
	tests := []struct {
		name string
		args args
		want *redditv1.SubredditPost
	}{
		{
			name: "convert",
			args: args{
				srp: datastore.SubredditPost{
					Author:         "fun_times",
					AuthorFullname: "aabbgghhss",
					Posts:          32,
				},
			},
			want: &redditv1.SubredditPost{
				Author:         "fun_times",
				AuthorFullname: "aabbgghhss",
				Posts:          32,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SubredditPostToProto(tt.args.srp)
			assert.Equal(t, tt.want, got)
		})
	}
}
