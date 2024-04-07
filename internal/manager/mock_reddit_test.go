// Code generated by MockGen. DO NOT EDIT.
// Source: internal/manager/reddit.go
//
// Generated by this command:
//
//	mockgen -source=internal/manager/reddit.go -destination=internal/manager/mock_reddit_test.go -package=manager
//

// Package manager is a generated GoMock package.
package manager

import (
	context "context"
	reflect "reflect"

	datastore "github.com/g8rswimmer/sub-reddit-stats/internal/datastore"
	gomock "go.uber.org/mock/gomock"
)

// MockFetcher is a mock of Fetcher interface.
type MockFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockFetcherMockRecorder
}

// MockFetcherMockRecorder is the mock recorder for MockFetcher.
type MockFetcherMockRecorder struct {
	mock *MockFetcher
}

// NewMockFetcher creates a new mock instance.
func NewMockFetcher(ctrl *gomock.Controller) *MockFetcher {
	mock := &MockFetcher{ctrl: ctrl}
	mock.recorder = &MockFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFetcher) EXPECT() *MockFetcherMockRecorder {
	return m.recorder
}

// SubredditPosts mocks base method.
func (m *MockFetcher) SubredditPosts(ctx context.Context, subreddit string, limit int) ([]datastore.SubredditPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubredditPosts", ctx, subreddit, limit)
	ret0, _ := ret[0].([]datastore.SubredditPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditPosts indicates an expected call of SubredditPosts.
func (mr *MockFetcherMockRecorder) SubredditPosts(ctx, subreddit, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditPosts", reflect.TypeOf((*MockFetcher)(nil).SubredditPosts), ctx, subreddit, limit)
}

// SubredditUps mocks base method.
func (m *MockFetcher) SubredditUps(ctx context.Context, subreddit string, limit int) ([]datastore.SubredditListing, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubredditUps", ctx, subreddit, limit)
	ret0, _ := ret[0].([]datastore.SubredditListing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditUps indicates an expected call of SubredditUps.
func (mr *MockFetcherMockRecorder) SubredditUps(ctx, subreddit, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditUps", reflect.TypeOf((*MockFetcher)(nil).SubredditUps), ctx, subreddit, limit)
}
