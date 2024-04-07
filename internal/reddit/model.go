package reddit

import (
	"time"
)

// Listing is the subreddit listing response.
// Rate limiting is include from the HTTP response
type Listing struct {
	Kind         string      `json:"kind"`
	Data         ListingData `json:"data"`
	RateLimiting *RateLimiting
}

// ListingData contains the subreddit data
type ListingData struct {
	After    string           `json:"after"`
	Before   string           `json:"before"`
	Children []SubredditChild `json:"children"`
}

// SubredditChild has the listing post data
type SubredditChild struct {
	Kind string        `json:"kind"`
	Data SubredditData `json:"data"`
}

// SubredditData has the post data
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
