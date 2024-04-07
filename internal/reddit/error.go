package reddit

import "fmt"

type HTTPError struct {
	StatusCode   int
	RateLimiting *RateLimiting
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
