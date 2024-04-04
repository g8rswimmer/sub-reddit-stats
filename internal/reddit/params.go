package reddit

import (
	"net/url"
	"strconv"
)

const (
	afterParam = "after"
	limitParam = "limit"
)

type Params func(url.Values)

func WithAfter(after string) Params {
	return func(v url.Values) {
		v.Add(afterParam, after)
	}
}

func WithLimit(limit int) Params {
	return func(v url.Values) {
		v.Add(limitParam, strconv.Itoa(limit))
	}
}
