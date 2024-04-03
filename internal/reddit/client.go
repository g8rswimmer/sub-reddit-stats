package reddit

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
)

const (
	subredditListingEndpoint = "/r/%s/new"
)

type Authorizer interface {
	AddAuthorization(req *http.Request)
}

type Client struct {
	BaseURL    string
	Auth       Authorizer
	HTTPClient *http.Client
}

func (c *Client) SubredditListingNow(ctx context.Context, subreddit string, params ...Params) (*model.RedditListing, error) {
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

	for k, v := range resp.Header {
		fmt.Println(k, v)
	}

	rateLimit := rateLimiting(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, &model.HTTPError{
			StatusCode:   resp.StatusCode,
			RateLimiting: rateLimit,
		}
	}

	listing := &model.RedditListing{}
	if err := json.NewDecoder(resp.Body).Decode(listing); err != nil {
		return nil, fmt.Errorf("subreddit listing response json decode: %w", err)
	}

	listing.RateLimiting = rateLimit

	return listing, nil
}
