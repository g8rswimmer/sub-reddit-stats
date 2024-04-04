package reddit

import (
	"context"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/g8rswimmer/sub-reddit-stats/internal/model"
	"github.com/stretchr/testify/assert"
)

type authMock struct{}

func (a authMock) AddAuthorization(req *http.Request) {}

func TestClient_SubredditListingNew(t *testing.T) {
	type fields struct {
		BaseURL    string
		Auth       Authorizer
		HTTPClient *http.Client
	}
	type args struct {
		subreddit string
		params    []Params
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.RedditListing
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				BaseURL: "https://www.test.com",
				Auth:    &authMock{},
				HTTPClient: mockHTTPClient(func(req *http.Request) *http.Response {
					if req.Method != http.MethodGet {
						log.Panicf("the method is not correct %s %s", req.Method, http.MethodGet)
					}
					body := `{
						"kind": "Listing",
						"data": {
							"after": "t3_1bv8ijk",
							"dist": 1,
							"modhash": "",
							"geo_filter": "",
							"children": [
								{
									"kind": "t3",
									"data": {
										"approved_at_utc": null,
										"subreddit": "funny",
										"selftext": "",
										"author_fullname": "t2_hfs25",
										"saved": false,
										"mod_reason_title": null,
										"gilded": 0,
										"clicked": false,
										"title": "Registering my kid for kindergarten...Do you think they'd honor it? 😂",
										"link_flair_richtext": [],
										"subreddit_name_prefixed": "r/funny",
										"hidden": false,
										"pwls": 6,
										"link_flair_css_class": null,
										"downs": 0,
										"thumbnail_height": 140,
										"top_awarded_type": null,
										"hide_score": true,
										"name": "t3_1bv8ijk",
										"quarantine": false,
										"link_flair_text_color": "dark",
										"upvote_ratio": 1.0,
										"author_flair_background_color": null,
										"subreddit_type": "public",
										"ups": 1,
										"total_awards_received": 0,
										"media_embed": {},
										"thumbnail_width": 140,
										"author_flair_template_id": null,
										"is_original_content": false,
										"user_reports": [],
										"secure_media": null,
										"is_reddit_media_domain": true,
										"is_meta": false,
										"category": null,
										"secure_media_embed": {},
										"link_flair_text": null,
										"can_mod_post": false,
										"score": 1,
										"approved_by": null,
										"is_created_from_ads_ui": false,
										"author_premium": false,
										"thumbnail": "https://a.thumbs.redditmedia.com/d0UNX-5-AQlo1DwavhbM3OonB85EzuDAH82ND1ys1K8.jpg",
										"edited": false,
										"author_flair_css_class": null,
										"author_flair_richtext": [],
										"gildings": {},
										"post_hint": "image",
										"content_categories": null,
										"is_self": false,
										"mod_note": null,
										"created": 1712188554.0,
										"link_flair_type": "text",
										"wls": 6,
										"removed_by_category": null,
										"banned_by": null,
										"author_flair_type": "text",
										"domain": "i.redd.it",
										"allow_live_comments": false,
										"selftext_html": null,
										"likes": null,
										"suggested_sort": null,
										"banned_at_utc": null,
										"url_overridden_by_dest": "https://i.redd.it/pckwpvikqcsc1.jpeg",
										"view_count": null,
										"archived": false,
										"no_follow": true,
										"is_crosspostable": false,
										"pinned": false,
										"over_18": false,
										"preview": {
											"images": [
												{
													"source": {
														"url": "https://preview.redd.it/pckwpvikqcsc1.jpeg?auto=webp&amp;s=42ad4e3eb425d18e41cff01099d1970cbbacf5a4",
														"width": 810,
														"height": 1080
													},
													"resolutions": [
														{
															"url": "https://preview.redd.it/pckwpvikqcsc1.jpeg?width=108&amp;crop=smart&amp;auto=webp&amp;s=4a9dd72f1f354a583a5f6442b8c9f33df6c17164",
															"width": 108,
															"height": 144
														},
														{
															"url": "https://preview.redd.it/pckwpvikqcsc1.jpeg?width=216&amp;crop=smart&amp;auto=webp&amp;s=e88352d0e4be9d83b71fd9e33f8050e4e85c9992",
															"width": 216,
															"height": 288
														},
														{
															"url": "https://preview.redd.it/pckwpvikqcsc1.jpeg?width=320&amp;crop=smart&amp;auto=webp&amp;s=83aa03770e6c884225509a20d65e4e36f3ecac6e",
															"width": 320,
															"height": 426
														},
														{
															"url": "https://preview.redd.it/pckwpvikqcsc1.jpeg?width=640&amp;crop=smart&amp;auto=webp&amp;s=b6e573aeb1990587e00a9300c1045fd5b8c712db",
															"width": 640,
															"height": 853
														}
													],
													"variants": {},
													"id": "FVVw5Iha4I5cRYoD6TrZnS6R_FCZw-QcsFOAGhqToYY"
												}
											],
											"enabled": true
										},
										"all_awardings": [],
										"awarders": [],
										"media_only": false,
										"can_gild": false,
										"spoiler": false,
										"locked": false,
										"author_flair_text": null,
										"treatment_tags": [],
										"visited": false,
										"removed_by": null,
										"num_reports": null,
										"distinguished": null,
										"subreddit_id": "t5_2qh33",
										"author_is_blocked": false,
										"mod_reason_by": null,
										"removal_reason": null,
										"link_flair_background_color": "",
										"id": "1bv8ijk",
										"is_robot_indexable": true,
										"report_reasons": null,
										"author": "dbzcat",
										"discussion_type": null,
										"num_comments": 1,
										"send_replies": true,
										"whitelist_status": "all_ads",
										"contest_mode": false,
										"mod_reports": [],
										"author_patreon_flair": false,
										"author_flair_text_color": null,
										"permalink": "/r/funny/comments/1bv8ijk/registering_my_kid_for_kindergartendo_you_think/",
										"parent_whitelist_status": "all_ads",
										"stickied": false,
										"url": "https://i.redd.it/pckwpvikqcsc1.jpeg",
										"subreddit_subscribers": 58221053,
										"created_utc": 1712188554.0,
										"num_crossposts": 0,
										"media": null,
										"is_video": false
									}
								}
							],
							"before": null
						}
					}`
					return &http.Response{
						StatusCode: http.StatusOK,
						Body:       io.NopCloser(strings.NewReader(body)),
						Header: func() http.Header {
							hdr := http.Header{}
							hdr.Add("x-ratelimit-remaining", "599.0")
							hdr.Add("x-ratelimit-used", "1")
							hdr.Add("x-ratelimit-reset", "569")
							return hdr
						}(),
					}
				}),
			},
			args: args{
				subreddit: "funny",
			},
			want: &model.RedditListing{
				Kind: "Listing",
				Data: model.RedditListingData{
					After: "t3_1bv8ijk",
					Children: []model.SubrredditChild{
						{
							Kind: "t3",
							Data: model.SubredditData{
								Title:               "Registering my kid for kindergarten...Do you think they'd honor it? 😂",
								Downs:               0,
								UpvoteRatio:         1.0,
								Ups:                 1,
								TotalAwardsReceived: 0,
								Name:                "t3_1bv8ijk",
								Subreddit:           "funny",
								ID:                  "1bv8ijk",
								Author:              "dbzcat",
							},
						},
					},
				},
				RateLimiting: &model.RateLimiting{
					Remaining: 599,
					Used:      1,
					Reset:     569 * time.Second,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				BaseURL:    tt.fields.BaseURL,
				Auth:       tt.fields.Auth,
				HTTPClient: tt.fields.HTTPClient,
			}
			got, err := c.SubredditListingNew(context.Background(), tt.args.subreddit, tt.args.params...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.SubredditListingNew() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
