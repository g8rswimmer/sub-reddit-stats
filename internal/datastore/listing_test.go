package datastore

import (
	"context"
	"sort"
	"testing"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestPresister_Store(t *testing.T) {
	type args struct {
		children []SubredditListing
	}
	tests := []struct {
		name    string
		args    args
		want    []SubredditListing
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				children: []SubredditListing{
					{
						ID:                  "t3_1bu4fzc",
						Downs:               1,
						UpvoteRatio:         0.5,
						Ups:                 2,
						TotalAwardsReceived: 10,
						Name:                "1bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
				},
			},
			want: []SubredditListing{
				{
					ID:                  "t3_1bu4fzc",
					Downs:               1,
					UpvoteRatio:         0.5,
					Ups:                 2,
					TotalAwardsReceived: 10,
					Name:                "1bu4fzc",
					Subreddit:           "funny",
					Author:              "sjustice",
					AuthorFullname:      "t2_bskdv",
				},
			},
			wantErr: false,
		},
		{
			name: "conflict",
			args: args{
				children: []SubredditListing{
					{
						ID:                  "t3_1bu4fzc",
						Downs:               1,
						UpvoteRatio:         0.5,
						Ups:                 2,
						TotalAwardsReceived: 10,
						Name:                "1bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_1bu4fzc",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 2000,
						TotalAwardsReceived: 10,
						Name:                "1bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
				},
			},
			want: []SubredditListing{
				{
					ID:                  "t3_1bu4fzc",
					Downs:               10,
					UpvoteRatio:         0.5,
					Ups:                 2000,
					TotalAwardsReceived: 10,
					Name:                "1bu4fzc",
					Subreddit:           "funny",
					Author:              "sjustice",
					AuthorFullname:      "t2_bskdv",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := DatabaseSetup()
			assert.NoError(t, err)
			defer db.Close()

			p := &Listing{
				DB: db,
			}

			if err := p.Store(context.Background(), tt.args.children); (err != nil) != tt.wantErr {
				t.Errorf("Presister.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()
			q := `SELECT * FROM listing`
			rows, err := db.QueryxContext(ctx, q)
			assert.NoError(t, err)
			defer rows.Close()

			expected := []SubredditListing{}

			for rows.Next() {
				data := SubredditListing{}
				err := rows.StructScan(&data)
				assert.NoError(t, err)
				expected = append(expected, data)
			}

			sort.Slice(expected, func(i, k int) bool {
				return expected[i].ID < expected[k].ID
			})
			sort.Slice(tt.want, func(i, k int) bool {
				return tt.want[i].ID < tt.want[k].ID
			})

			assert.Equal(t, tt.want, expected)

		})
	}
}

func TestPresister_SubredditUps(t *testing.T) {
	type seed struct {
		children []SubredditListing
	}
	type args struct {
		subreddit string
		limit     int
	}
	tests := []struct {
		name    string
		seed    seed
		args    args
		want    []model.SubredditData
		wantErr bool
	}{
		{
			name: "success",
			seed: seed{
				children: []SubredditListing{
					{
						ID:                  "t3_1bu4fzc",
						Downs:               1,
						UpvoteRatio:         0.5,
						Ups:                 2,
						TotalAwardsReceived: 10,
						Name:                "1bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_2bu4fzc",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 2000,
						TotalAwardsReceived: 10,
						Name:                "2bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_3bu4fzc",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 10,
						TotalAwardsReceived: 10,
						Name:                "3bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
				},
			},
			args: args{
				subreddit: "funny",
				limit:     5,
			},
			want: []model.SubredditData{
				{
					ID:                  "t3_2bu4fzc",
					Downs:               10,
					UpvoteRatio:         0.5,
					Ups:                 2000,
					TotalAwardsReceived: 10,
					Name:                "2bu4fzc",
					Subreddit:           "funny",
					Author:              "sjustice",
					AuthorFullname:      "t2_bskdv",
				},
				{
					ID:                  "t3_3bu4fzc",
					Downs:               10,
					UpvoteRatio:         0.5,
					Ups:                 10,
					TotalAwardsReceived: 10,
					Name:                "3bu4fzc",
					Subreddit:           "funny",
					Author:              "sjustice",
					AuthorFullname:      "t2_bskdv",
				},
				{
					ID:                  "t3_1bu4fzc",
					Downs:               1,
					UpvoteRatio:         0.5,
					Ups:                 2,
					TotalAwardsReceived: 10,
					Name:                "1bu4fzc",
					Subreddit:           "funny",
					Author:              "sjustice",
					AuthorFullname:      "t2_bskdv",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := DatabaseSetup()
			assert.NoError(t, err)
			defer db.Close()

			p := &Listing{
				DB: db,
			}

			if err := p.Store(context.Background(), tt.seed.children); (err != nil) != tt.wantErr {
				t.Errorf("Presister.SubredditUps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := p.SubredditUps(context.Background(), tt.args.subreddit, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Presister.SubredditUps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestListing_SubredditPosts(t *testing.T) {
	type seed struct {
		children []SubredditListing
	}
	type args struct {
		subreddit string
		limit     int
	}
	tests := []struct {
		name    string
		seed    seed
		args    args
		want    []model.SubredditPost
		wantErr bool
	}{
		{
			name: "success",
			seed: seed{
				children: []SubredditListing{
					{
						ID:                  "t3_1bu4fza",
						Downs:               1,
						UpvoteRatio:         0.5,
						Ups:                 2,
						TotalAwardsReceived: 10,
						Name:                "1bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_2bu4fzb",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 2000,
						TotalAwardsReceived: 10,
						Name:                "2bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_3bu4fzc",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 10,
						TotalAwardsReceived: 10,
						Name:                "3bu4fzc",
						Subreddit:           "funny",
						Author:              "sjustice",
						AuthorFullname:      "t2_bskdv",
					},
					{
						ID:                  "t3_3bu4faa",
						Downs:               10,
						UpvoteRatio:         0.5,
						Ups:                 10,
						TotalAwardsReceived: 10,
						Name:                "3bu4fzc",
						Subreddit:           "funny",
						Author:              "rjustice",
						AuthorFullname:      "t2_bskdz",
					},
				},
			},
			args: args{
				subreddit: "funny",
				limit:     1,
			},
			want: []model.SubredditPost{
				{
					Author:         "sjustice",
					AuthorFullname: "t2_bskdv",
					Posts:          3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := DatabaseSetup()
			assert.NoError(t, err)
			defer db.Close()

			p := &Listing{
				DB: db,
			}
			if err := p.Store(context.Background(), tt.seed.children); (err != nil) != tt.wantErr {
				t.Errorf("Presister.SubredditPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := p.SubredditPosts(context.Background(), tt.args.subreddit, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("Listing.SubredditPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
