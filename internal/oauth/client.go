package oauth

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const (
	endpoint  = "/api/v1/access_token"
	grantType = "client_credentials"
)

type client struct {
	httpClient   *http.Client
	baseURL      string
	clientID     string
	clientSecret string
	deviceID     string
}

// AccessToken will retreve an access token from the reddit oauth server
func (c *client) AccessToken(ctx context.Context) (*accessResponse, error) {
	data := url.Values{}
	data.Set("grant_type", grantType)
	data.Set("device_id", c.deviceID)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("oauth request err: %w", err)
	}
	req.SetBasicAuth(c.clientID, c.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("oauth response err: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("oauth response code %d %s", resp.StatusCode, resp.Status)
	}
	ar := &accessResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ar); err != nil {
		return nil, fmt.Errorf("oauth response decode err: %w", err)
	}
	return ar, nil
}
