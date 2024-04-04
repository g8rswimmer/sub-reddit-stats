package reddit

import (
	"net/url"
	"strconv"
)

const (
	beforeParam = "before"
	limitParam  = "limit"
)

type Params func(url.Values)

func WithBefore(before string) Params {
	return func(v url.Values) {
		v.Add(beforeParam, before)
	}
}

func WithLimit(limit int) Params {
	return func(v url.Values) {
		v.Add(limitParam, strconv.Itoa(limit))
	}
}
