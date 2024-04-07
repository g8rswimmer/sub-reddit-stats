package worker

import (
	"testing"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	model "github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestRunner_process(t *testing.T) {
	type fields struct {
		Lister    func(ctrl *gomock.Controller) RedditLister
		Presister func(ctrl *gomock.Controller) Presister
	}
	type args struct {
		subreddit string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    time.Duration
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Lister: func(ctrl *gomock.Controller) RedditLister {
					m := NewMockRedditLister(ctrl)
					r := &model.RedditListing{
						Kind: "Listing",
						Data: model.RedditListingData{
							After: "t3_1bv8ijk",
							Children: []model.SubredditChild{
								{
									Kind: "t3",
									Data: model.SubredditData{
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
							},
						},
						RateLimiting: &model.RateLimiting{
							Remaining: 1,
							Used:      1,
							Reset:     569 * time.Second,
						},
					}
					m.EXPECT().SubredditListingNew(gomock.All(), "funny", gomock.Any()).Return(r, nil)
					return m
				},
				Presister: func(ctrl *gomock.Controller) Presister {
					m := NewMockPresister(ctrl)
					children := []datastore.SubredditListing{
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
					m.EXPECT().Store(gomock.Any(), children).Return(nil)
					return m
				},
			},
			args: args{
				subreddit: "funny",
			},
			want: 569*time.Second + time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			r := &Runner{
				Lister:    tt.fields.Lister(ctrl),
				Presister: tt.fields.Presister(ctrl),
			}
			got, err := r.process(tt.args.subreddit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Runner.process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
