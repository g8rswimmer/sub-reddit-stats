package oauth

import (
	"net/http"
	"os"
	"time"
)

const (
	baseURL         = "https://www.reddit.com"
	clientIDKey     = "CLIENT_ID"
	clientSecretKey = "CLIENT_SECRET"
)

// Option are oauth manager configuratoin options
type Option func(*config)

type config struct {
	baseURL      string
	clientID     string
	clientSecret string
	httpClient   *http.Client
	deviceID     string
}

func loadDefaultConfig() *config {
	return &config{
		baseURL:      baseURL,
		clientID:     os.Getenv(clientIDKey),
		clientSecret: os.Getenv(clientSecretKey),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		deviceID: "a56c8f203485e862b22ac55ad0a05",
	}
}

// WithBaseURL will set the base URL for the oauth manager.  This is
// used to retrieve tokens
func WithBaseURL(baseURL string) Option {
	return func(cfg *config) {
		cfg.baseURL = baseURL
	}
}

// WithCredentials the client ID and secret that will be used to
// obtain access tokens
func WithCredentials(clientID, clientSecret string) Option {
	return func(cfg *config) {
		cfg.clientID = clientID
		cfg.clientSecret = clientSecret
	}
}

// WithHTTPClient will set the HTTP client that will be used to
// make token callouts
func WithHTTPClient(client *http.Client) Option {
	return func(cfg *config) {
		cfg.httpClient = client
	}
}

// WithDeviceID is the randome device ID used when calling
// the oauth endpoint
func WithDeviceID(id string) Option {
	return func(cfg *config) {
		cfg.deviceID = id
	}
}
