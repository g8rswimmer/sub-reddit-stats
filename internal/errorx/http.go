package errorx

import (
	"fmt"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
)

type HTTPError struct {
	StatusCode   int
	RateLimiting *model.RateLimiting
}

func (h *HTTPError) Error() string {
	return fmt.Sprintf("http error response %d", h.StatusCode)
}

func (h *HTTPError) Is(target error) bool {
	cmp, ok := target.(*HTTPError)
	if !ok {
		return false
	}
	return cmp.StatusCode == h.StatusCode
}
