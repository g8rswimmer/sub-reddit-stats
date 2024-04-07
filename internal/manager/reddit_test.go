package manager

import (
	"context"
	"testing"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestReddit_SubredditMostUps(t *testing.T) {
	type fields struct {
		Fetcher func(ctrl *gomock.Controller) Fetcher
	}
	type args struct {
		subreddit string
		limit     int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*redditv1.SubredditData
		wantErr bool
	}{
		{
			name: "simple",
			fields: fields{
				Fetcher: func(ctrl *gomock.Controller) Fetcher {
					m := NewMockFetcher(ctrl)
					r := []model.SubredditData{
						{
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
					}
					m.EXPECT().SubredditUps(gomock.Any(), "funny", 1).Return(r, nil)
					return m
				},
			},
			args: args{
				subreddit: "funny",
				limit:     1,
			},
			want: []*redditv1.SubredditData{
				{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			r := &Reddit{
				Fetcher: tt.fields.Fetcher(ctrl),
			}
			got, err := r.SubredditMostUps(context.Background(), tt.args.subreddit, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reddit.SubredditMostUps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}