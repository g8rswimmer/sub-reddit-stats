package model

import "time"

type RedditListing struct {
	Kind         string            `json:"kind"`
	Data         RedditListingData `json:"data"`
	RateLimiting *RateLimiting
}

type RedditListingData struct {
	After    string           `json:"after"`
	Before   string           `json:"before"`
	Children []SubredditChild `json:"children"`
}

type SubredditChild struct {
	Kind string        `json:"kind"`
	Data SubredditData `json:"data"`
}

type SubredditData struct {
	Title               string  `json:"title" db:"TITLE"`
	Downs               int     `json:"downs" db:"DOWNS"`
	UpvoteRatio         float64 `json:"upvote_ratio" db:"UPVOTE_RATIO"`
	Ups                 int     `json:"ups" db:"UPS"`
	TotalAwardsReceived int     `json:"total_awards_received" db:"TOTAL_AWARDS"`
	Name                string  `json:"name" db:"NAME"`
	Subreddit           string  `json:"subreddit" db:"SUBREDDIT"`
	ID                  string  `json:"id" db:"ID"`
	Author              string  `json:"author" db:"AUTHOR"`
	AuthorFullname      string  `json:"author_fullname" db:"AUTHOR_FULLNAME"`
}

type RateLimiting struct {
	Remaining int
	Used      int
	Reset     time.Duration
}

type SubredditPost struct {
	Author         string `json:"author" db:"AUTHOR"`
	AuthorFullname string `json:"author_fullname" db:"AUTHOR_FULLNAME"`
	Posts          int    `json:"posts" db:"POSTS"`
}
