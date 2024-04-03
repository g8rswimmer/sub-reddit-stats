package reddit

import (
	"net/http"
	"testing"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/stretchr/testify/assert"
)

func Test_rateLimiting(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name string
		args args
		want *model.RateLimiting
	}{
		{
			name: "success",
			args: args{
				resp: &http.Response{
					Header: func() http.Header {
						h := http.Header{}
						h.Add("x-ratelimit-remaining", "599.0")
						h.Add("x-ratelimit-used", "1")
						h.Add("x-ratelimit-reset", "247")
						return h
					}(),
				},
			},
			want: &model.RateLimiting{
				Remaining: 599,
				Used:      1,
				Reset:     247 * time.Second,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rateLimiting(tt.args.resp)

			assert.Equal(t, tt.want, got)
		})
	}
}
