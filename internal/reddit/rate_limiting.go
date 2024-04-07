package reddit

import (
	"net/http"
	"strconv"
	"time"
)

const (
	rateLimitingRemaining = "x-ratelimit-remaining"
	rateLimitingUsed      = "x-ratelimit-used"
	rateLimitingReset     = "x-ratelimit-reset"
)

func rateLimiting(resp *http.Response) *RateLimiting {
	return &RateLimiting{
		Remaining: func() int {
			r := resp.Header.Get(rateLimitingRemaining)
			if len(r) == 0 {
				return 0
			}
			remaining, err := strconv.ParseFloat(r, 64)
			if err != nil {
				return 0
			}
			return int(remaining)
		}(),
		Used: func() int {
			r := resp.Header.Get(rateLimitingUsed)
			if len(r) == 0 {
				return 0
			}
			used, err := strconv.ParseFloat(r, 64)
			if err != nil {
				return 0
			}
			return int(used)
		}(),
		Reset: func() time.Duration {
			r := resp.Header.Get(rateLimitingReset)
			if len(r) == 0 {
				return 0
			}
			reset, err := strconv.ParseFloat(r, 64)
			if err != nil {
				return 0
			}
			return time.Duration(reset) * time.Second
		}(),
	}
}
