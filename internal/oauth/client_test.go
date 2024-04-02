package oauth

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_client_AccessToken(t *testing.T) {
	type fields struct {
		httpClient   *http.Client
		baseURL      string
		clientID     string
		clientSecret string
		deviceID     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *accessResponse
		wantErr bool
	}{
		{
			name: "token response success",
			fields: fields{
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
			want: &accessResponse{
				AccessToken: "this is the access token",
				TokenType:   "bearer",
				ExpiresIn:   3600,
				Scope:       "*",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				httpClient:   tt.fields.httpClient,
				baseURL:      tt.fields.baseURL,
				clientID:     tt.fields.clientID,
				clientSecret: tt.fields.clientSecret,
				deviceID:     tt.fields.deviceID,
			}
			got, err := c.AccessToken(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("client.AccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
