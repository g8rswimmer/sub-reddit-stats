package oauth

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestManager_initAccessToken(t *testing.T) {
	type fields struct {
		accessClient *client
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				accessClient: &client{
					httpClient: mockHTTPClient(func(req *http.Request) *http.Response {
						if req.Method != http.MethodPost {
							log.Panicf("the method is not correct %s %s", req.Method, http.MethodGet)
						}
						body := `{
							"access_token": "this is the access token",
							"token_type": "bearer",
							"expires_in": 3600,
							"scope": "*"
						}`
						return &http.Response{
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(strings.NewReader(body)),
						}
					}),
					baseURL:      "http://www.test.com",
					clientID:     "id",
					clientSecret: "shhhhh",
					deviceID:     "11223344",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				accessClient: tt.fields.accessClient,
			}
			if err := m.initAccessToken(context.Background()); (err != nil) != tt.wantErr {
				t.Errorf("Manager.initAccessToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestManager_setToken(t *testing.T) {
	type args struct {
		ar *accessResponse
	}
	tests := []struct {
		name string
		args args
		want time.Duration
	}{
		{
			name: "success",
			args: args{
				ar: &accessResponse{
					AccessToken: "this is the access token",
					TokenType:   "bearer",
					ExpiresIn:   3600,
					Scope:       "*",
				},
			},
			want: 1800 * time.Second,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{}
			if got := m.setToken(tt.args.ar); got != tt.want {
				t.Errorf("Manager.setToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestManager_AddAuthorization(t *testing.T) {
	type fields struct {
		accessResponse *accessResponse
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Success",
			fields: fields{
				accessResponse: &accessResponse{
					AccessToken: "this is the access token",
					TokenType:   "bearer",
					ExpiresIn:   3600,
					Scope:       "*",
				},
			},
			args: args{
				req: httptest.NewRequest(http.MethodPost, "https://www.google.com/hello", nil),
			},
			want: "bearer this is the access token",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Manager{
				accessResponse: tt.fields.accessResponse,
			}
			m.AddAuthorization(tt.args.req)
			assert.Equal(t, tt.want, tt.args.req.Header.Get("Authorization"))

		})
	}
}
