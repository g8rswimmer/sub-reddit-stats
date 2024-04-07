package service

import (
	"context"
	"testing"

	"github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	"github.com/google/go-cmp/cmp"
	gomock "go.uber.org/mock/gomock"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestReddit_GetSubredditMostUps(t *testing.T) {
	type fields struct {
		Manager func(ctrl *gomock.Controller) Manager
	}
	type args struct {
		req *redditv1.GetSubredditMostUpsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *redditv1.GetSubredditMostUpsResponse
		wantErr bool
	}{
		{
			name: "simple success",
			fields: fields{
				Manager: func(ctrl *gomock.Controller) Manager {
					m := NewMockManager(ctrl)
					r := []*redditv1.SubredditData{
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
					}
					m.EXPECT().SubredditMostUps(gomock.Any(), "funny", 1).Return(r, nil)
					return m
				},
			},
			args: args{
				req: &redditv1.GetSubredditMostUpsRequest{
					Subreddit: "funny",
					Limit:     1,
				},
			},
			want: &redditv1.GetSubredditMostUpsResponse{
				SubredditPosts: []*redditv1.SubredditData{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			r := &Reddit{
				Manager: tt.fields.Manager(ctrl),
			}
			got, err := r.GetSubredditMostUps(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reddit.GetSubredditMostUps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				protocmp.Transform(),
			}
			if !cmp.Equal(got, tt.want, opts) {
				t.Errorf("Reddit.GetSubredditMostUps() (-want, +got) %s", cmp.Diff(tt.want, got, opts))
			}
		})
	}
}

func TestReddit_GetSubredditAuthorPosts(t *testing.T) {
	type fields struct {
		Manager func(ctrl *gomock.Controller) Manager
	}
	type args struct {
		req *redditv1.GetSubredditAuthorPostsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *redditv1.GetSubredditAuthorPostsResponse
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				Manager: func(ctrl *gomock.Controller) Manager {
					m := NewMockManager(ctrl)
					r := []*redditv1.SubredditPost{
						{
							Author:         "dbzcat",
							AuthorFullname: "aabbgg",
							Posts:          34,
						},
					}
					m.EXPECT().SubredditAuthorPosts(gomock.Any(), "funny", 1).Return(r, nil)
					return m
				},
			},
			args: args{
				req: &redditv1.GetSubredditAuthorPostsRequest{
					Subreddit: "funny",
					Limit:     1,
				},
			},
			want: &redditv1.GetSubredditAuthorPostsResponse{
				AuthorPosts: []*redditv1.SubredditPost{
					{
						Author:         "dbzcat",
						AuthorFullname: "aabbgg",
						Posts:          34,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)

			r := &Reddit{
				Manager: tt.fields.Manager(ctrl),
			}
			got, err := r.GetSubredditAuthorPosts(context.Background(), tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reddit.GetSubredditAuthorPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opts := cmp.Options{
				protocmp.Transform(),
			}
			if !cmp.Equal(got, tt.want, opts) {
				t.Errorf("Reddit.GetSubredditMostUps() (-want, +got) %s", cmp.Diff(tt.want, got, opts))
			}
		})
	}
}
