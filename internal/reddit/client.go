package reddit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	subredditListingEndpoint = "/r/%s/new"
)

// Authorizer will set the authrorization headers in the HTTP request
type Authorizer interface {
	AddAuthorization(req *http.Request)
}

// Client is used to interface with the reddit APIs.
type Client struct {
	BaseURL    string
	Auth       Authorizer
	HTTPClient *http.Client
}

// SubredditListingNew will retrieve the subreddit listings.  This is from the reddit APIs with documentation.
// https://www.reddit.com/dev/api/#GET_new
func (c *Client) SubredditListingNew(ctx context.Context, subreddit string, params ...Params) (*Listing, error) {
	ep := fmt.Sprintf(subredditListingEndpoint, subreddit)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+ep, nil)
	if err != nil {
		return nil, fmt.Errorf("subreddit listing request: %w", err)
	}
	q := req.URL.Query()
	for _, param := range params {
		param(q)
	}
	req.URL.RawQuery = q.Encode()

	c.Auth.AddAuthorization(req)
	req.Header.Add("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("subreddit listing response: %w", err)
	}
	defer resp.Body.Close()

	rateLimit := rateLimiting(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, &HTTPError{
			StatusCode:   resp.StatusCode,
			RateLimiting: rateLimit,
		}
	}

	listing := &Listing{}
	if err := json.NewDecoder(resp.Body).Decode(listing); err != nil {
		return nil, fmt.Errorf("subreddit listing response json decode: %w", err)
	}
	listing.RateLimiting = rateLimit

	return listing, nil
}
