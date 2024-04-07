package datastore

// SubredditListing is the database model for the subreddit listing
type SubredditListing struct {
	Title               string  `db:"TITLE"`
	Downs               int     `db:"DOWNS"`
	UpvoteRatio         float64 `db:"UPVOTE_RATIO"`
	Ups                 int     `db:"UPS"`
	TotalAwardsReceived int     `db:"TOTAL_AWARDS"`
	Name                string  `db:"NAME"`
	Subreddit           string  `db:"SUBREDDIT"`
	ID                  string  `db:"ID"`
	Author              string  `db:"AUTHOR"`
	AuthorFullname      string  `db:"AUTHOR_FULLNAME"`
}

// SubredditPost is the database model for the author's subreddit posts
type SubredditPost struct {
	Author         string `json:"author" db:"AUTHOR"`
	AuthorFullname string `json:"author_fullname" db:"AUTHOR_FULLNAME"`
	Posts          int    `json:"posts" db:"POSTS"`
}
