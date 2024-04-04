// Code generated by MockGen. DO NOT EDIT.
// Source: internal/worker/runner.go
//
// Generated by this command:
//
//	mockgen -source=internal/worker/runner.go -destination=internal/worker/mock_runner_test.go -package=worker
//

// Package worker is a generated GoMock package.
package worker

import (
	context "context"
	reflect "reflect"

	model "github.com/g8rswimmer/sub-reddit-stats/internal/model"
	reddit "github.com/g8rswimmer/sub-reddit-stats/internal/reddit"
	gomock "go.uber.org/mock/gomock"
)

// MockPresister is a mock of Presister interface.
type MockPresister struct {
	ctrl     *gomock.Controller
	recorder *MockPresisterMockRecorder
}

// MockPresisterMockRecorder is the mock recorder for MockPresister.
type MockPresisterMockRecorder struct {
	mock *MockPresister
}

// NewMockPresister creates a new mock instance.
func NewMockPresister(ctrl *gomock.Controller) *MockPresister {
	mock := &MockPresister{ctrl: ctrl}
	mock.recorder = &MockPresisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPresister) EXPECT() *MockPresisterMockRecorder {
	return m.recorder
}

// StoreListing mocks base method.
func (m *MockPresister) StoreListing(ctx context.Context, children []model.SubredditChild) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreListing", ctx, children)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreListing indicates an expected call of StoreListing.
func (mr *MockPresisterMockRecorder) StoreListing(ctx, children any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreListing", reflect.TypeOf((*MockPresister)(nil).StoreListing), ctx, children)
}

// MockRedditLister is a mock of RedditLister interface.
type MockRedditLister struct {
	ctrl     *gomock.Controller
	recorder *MockRedditListerMockRecorder
}

// MockRedditListerMockRecorder is the mock recorder for MockRedditLister.
type MockRedditListerMockRecorder struct {
	mock *MockRedditLister
}

// NewMockRedditLister creates a new mock instance.
func NewMockRedditLister(ctrl *gomock.Controller) *MockRedditLister {
	mock := &MockRedditLister{ctrl: ctrl}
	mock.recorder = &MockRedditListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedditLister) EXPECT() *MockRedditListerMockRecorder {
	return m.recorder
}

// SubredditListingNew mocks base method.
func (m *MockRedditLister) SubredditListingNew(ctx context.Context, subreddit string, params ...reddit.Params) (*model.RedditListing, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, subreddit}
	for _, a := range params {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubredditListingNew", varargs...)
	ret0, _ := ret[0].(*model.RedditListing)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubredditListingNew indicates an expected call of SubredditListingNew.
func (mr *MockRedditListerMockRecorder) SubredditListingNew(ctx, subreddit any, params ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, subreddit}, params...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubredditListingNew", reflect.TypeOf((*MockRedditLister)(nil).SubredditListingNew), varargs...)
}
