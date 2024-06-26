// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/reddit.go
//
// Generated by this command:
//
//	mockgen -source=internal/service/reddit.go -destination=internal/service/mock_reddit_test.go -package=service
//

// Package service is a generated GoMock package.
package service

import (
	context "context"
	reflect "reflect"

	redditv1 "github.com/g8rswimmer/sub-reddit-stats/internal/proto/redditv1"
	gomock "go.uber.org/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// SubredditAuthorPosts mocks base method.
func (m *MockManager) SubredditAuthorPosts(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditPost, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubredditAuthorPosts", ctx, subreddit, limit)
	ret0, _ := ret[0].([]*redditv1.SubredditPost)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditAuthorPosts indicates an expected call of SubredditAuthorPosts.
func (mr *MockManagerMockRecorder) SubredditAuthorPosts(ctx, subreddit, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditAuthorPosts", reflect.TypeOf((*MockManager)(nil).SubredditAuthorPosts), ctx, subreddit, limit)
}

// SubredditMostUps mocks base method.
func (m *MockManager) SubredditMostUps(ctx context.Context, subreddit string, limit int) ([]*redditv1.SubredditData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubredditMostUps", ctx, subreddit, limit)
	ret0, _ := ret[0].([]*redditv1.SubredditData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditMostUps indicates an expected call of SubredditMostUps.
func (mr *MockManagerMockRecorder) SubredditMostUps(ctx, subreddit, limit any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditMostUps", reflect.TypeOf((*MockManager)(nil).SubredditMostUps), ctx, subreddit, limit)
}
