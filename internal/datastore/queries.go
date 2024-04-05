package datastore

const insertListing = `INSERT INTO listing (
	"ID",
	"DOWNS",
	"UPS",
	"TITLE",
	"NAME",
	"SUBREDDIT",
	"AUTHOR",
	"AUTHOR_FULLNAME",
	"TOTAL_AWARDS",
	"UPVOTE_RATIO")
	VALUES (
		:ID,
		:DOWNS,
		:UPS,
		:TITLE,
		:NAME,
		:SUBREDDIT,
		:AUTHOR,
		:AUTHOR_FULLNAME,
		:TOTAL_AWARDS,
		:UPVOTE_RATIO)
	ON CONFLICT (ID) DO UPDATE SET
		DOWNS=excluded.DOWNS,
		UPS=excluded.UPS,
		TITLE=excluded.TITLE,
		NAME=excluded.NAME,
		SUBREDDIT=excluded.SUBREDDIT,
		AUTHOR=excluded.AUTHOR,
		AUTHOR_FULLNAME=excluded.AUTHOR_FULLNAME,
		TOTAL_AWARDS=excluded.TOTAL_AWARDS,
		UPVOTE_RATIO=excluded.UPVOTE_RATIO
		`

const listingUps = `SELECT * FROM listing WHERE SUBREDDIT = :subreddit ORDER BY UPS DESC LIMIT :limit`
