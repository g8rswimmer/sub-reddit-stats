package datastore

import (
	"context"
	"log/slog"
	"sort"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
)

type Presister struct{}

func (p *Presister) StoreListing(ctx context.Context, children []model.SubredditChild) error {
	sort.Slice(children, func(i, k int) bool {
		return children[i].Data.Ups > children[k].Data.Ups
	})
	end := len(children)
	if end > 5 {
		end = 5
	}
	for _, child := range children[:end] {
		slog.Info(child.Data.Title, "author", child.Data.Author, "ups", child.Data.Ups)
	}
	return nil
}
