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

func WithBaseURL(baseURL string) Option {
	return func(cfg *config) {
		cfg.baseURL = baseURL
	}
}

func WithCredentials(clientID, clientSecret string) Option {
	return func(cfg *config) {
		cfg.clientID = clientID
		cfg.clientSecret = clientSecret
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(cfg *config) {
		cfg.httpClient = client
	}
}

func WithDeviceID(id string) Option {
	return func(cfg *config) {
		cfg.deviceID = id
	}
}
