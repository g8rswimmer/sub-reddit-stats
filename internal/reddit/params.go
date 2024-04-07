package reddit

import (
	"net/url"
	"strconv"
)

const (
	afterParam = "after"
	limitParam = "limit"
)

// Params is used to add optional data to the request callout
type Params func(url.Values)

// WithAfter will set the after parameter for the API callout
func WithAfter(after string) Params {
	return func(v url.Values) {
		v.Add(afterParam, after)
	}
}

// WithLimit will set the limit paramter for the API callout
func WithLimit(limit int) Params {
	return func(v url.Values) {
		v.Add(limitParam, strconv.Itoa(limit))
	}
}
