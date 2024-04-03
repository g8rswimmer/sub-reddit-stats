package oauth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

const (
	tick    = 100 * time.Millisecond
	timeout = 2 * time.Second
)

type Manager struct {
	accessClient   *client
	mutex          sync.RWMutex
	accessResponse *accessResponse
	done           chan struct{}
}

func NewManager(ctx context.Context, options ...Option) (*Manager, error) {
	cfg := loadDefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}

	mngr := &Manager{
		accessClient: &client{
			baseURL:      cfg.baseURL,
			clientID:     cfg.clientID,
			clientSecret: cfg.clientSecret,
			deviceID:     cfg.deviceID,
			httpClient:   cfg.httpClient,
		},
		done: make(chan struct{}),
	}

	if err := mngr.initAccessToken(ctx); err != nil {
		return nil, err
	}
	mngr.run()
	slog.InfoContext(ctx, "oauth manager starting....")
	return mngr, nil
}

func (m *Manager) initAccessToken(ctx context.Context) error {
	ticker := time.NewTicker(tick)
	defer ticker.Stop()
	to := time.NewTimer(timeout)
	defer to.Stop()

	for {
		select {
		case <-ticker.C:
			ar, err := m.accessClient.AccessToken(ctx)
			if err == nil {
				m.accessResponse = ar
				return nil
			}
		case <-to.C:
			return errors.New("oauth manager token timeout")
		}
	}
}

func (m *Manager) run() {
	exp := time.Duration(m.accessResponse.ExpiresIn) * time.Second
	exp /= 2
	refresh := time.NewTimer(exp)

	go func() {
		select {
		case <-refresh.C:
			slog.Info("oauth manager access token refresh....")
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			ar, err := m.accessClient.AccessToken(ctx)
			cancel()
			switch {
			case err == nil:
				exp = m.setToken(ar)
				refresh.Reset(exp)
				slog.Info("oauth manager token set", "expire", exp.String())
			default:
				slog.Error("oauth manager access token", "error", err.Error())
				refresh.Reset(timeout)
			}
		case <-m.done:
			slog.Info("oauth manager shutting down")
			refresh.Stop()
			return
		}
	}()
}

func (m *Manager) setToken(ar *accessResponse) time.Duration {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.accessResponse = ar
	exp := time.Duration(m.accessResponse.ExpiresIn) * time.Second
	exp /= 2
	return exp
}

func (m *Manager) AddAuthorization(req *http.Request) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	auth := fmt.Sprintf("Bearer %s", m.accessResponse.AccessToken)
	req.Header.Add("Authorization", auth)
}

func (m *Manager) Shutdown() {
	close(m.done)
}
