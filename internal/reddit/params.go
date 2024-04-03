package reddit

import "net/url"

const (
	beforeParam = "before"
)

type Params func(url.Values)

func WithBefore(before string) Params {
	return func(v url.Values) {
		v.Add(beforeParam, before)
	}
}
