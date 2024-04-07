package reddit

import (
	"time"
)

type Listing struct {
	Kind         string      `json:"kind"`
	Data         ListingData `json:"data"`
	RateLimiting *RateLimiting
}

type ListingData struct {
	After    string           `json:"after"`
	Before   string           `json:"before"`
	Children []SubredditChild `json:"children"`
}

type SubredditChild struct {
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
	AuthorFullname      string  `json:"author_fullname"`
}

type RateLimiting struct {
	Remaining int
	Used      int
	Reset     time.Duration
}
