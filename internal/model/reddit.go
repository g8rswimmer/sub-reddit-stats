package model

import "time"

type RedditListing struct {
	Kind         string            `json:"kind"`
	Data         RedditListingData `json:"data"`
	RateLimiting *RateLimiting
}

type RedditListingData struct {
	After    string            `json:"after"`
	Before   string            `json:"before"`
	Children []SubrredditChild `json:"children"`
}

type SubrredditChild struct {
	Kind string        `json:"kind"`
	Data SubredditData `json:"data"`
}

type SubredditData struct {
	Title               string  `json:"title"`
	Downs               int     `json:"downs"`
	UpvoteRatio         float64 `json:"upvote_ratio"`
	Ups                 int     `json:"ups"`
	TotalAwardsReceived int     `json:"total_awards_received"`
	Name                string  `json:"name"`
	Subreddit           string  `json:"subreddit"`
	ID                  string  `json:"id"`
	Author              string  `json:"author"`
}

type RateLimiting struct {
	Remaining int
	Used      int
	Reset     time.Duration
}
