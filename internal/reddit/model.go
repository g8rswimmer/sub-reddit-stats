package reddit

import "time"

type RateLimiting struct {
	Remaining int
	Used      int
	Reset     time.Duration
}
