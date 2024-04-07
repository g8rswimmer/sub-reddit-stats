package datastore

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
