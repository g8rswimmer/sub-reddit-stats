package datastore

import (
	"context"
	"database/sql"
	"fmt"
)

const (
	listingTbl = `CREATE TABLE IF NOT EXISTS "listing" (
		"ID" VARCHAR(50) PRIMARY KEY,
		"DOWNS" INT,
		"UPS" INT,
		"TITLE" TEXT,
		"NAME" VARCAHR(50),
		"SUBREDDIT" VARCHAR(50),
		"AUTHOR" VARCAHR(200),
		"AUTHOR_FULLNAME" VARCHAR(50),
		"TOTAL_AWARDS" INT,
		"UPVOTE_RATIO" FLOAT
	)`

	authorIdx         = `CREATE INDEX IF NOT EXISTS author_idx ON listing (AUTHOR);`
	subredditIdx      = `CREATE INDEX IF NOT EXISTS subreddit_idx ON listing (SUBREDDIT);`
	authorFullNameIdx = `CREATE INDEX IF NOT EXISTS author_fullname_idx ON listing (AUTHOR_FULLNAME);`
)

type Migration struct {
	DB *sql.DB
}

func (m *Migration) Apply(ctx context.Context) error {
	if _, err := m.DB.ExecContext(ctx, listingTbl); err != nil {
		return fmt.Errorf("table creation error: %w", err)
	}
	if _, err := m.DB.ExecContext(ctx, authorIdx); err != nil {
		return fmt.Errorf("table creation error: %w", err)
	}
	if _, err := m.DB.ExecContext(ctx, authorFullNameIdx); err != nil {
		return fmt.Errorf("table creation error: %w", err)
	}
	return nil
}
