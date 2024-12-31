// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/diSpector/incrowd.git/internal/models/domain"
	mock "github.com/stretchr/testify/mock"
)

// PollStorage is an autogenerated mock type for the PollStorage type
type PollStorage struct {
	mock.Mock
}

// GetAtricleBySource provides a mock function with given fields: ctx, sourceId, sourceSystem
func (_m *PollStorage) GetAtricleBySource(ctx context.Context, sourceId string, sourceSystem string) (*domain.Article, error) {
	ret := _m.Called(ctx, sourceId, sourceSystem)

	if len(ret) == 0 {
		panic("no return value specified for GetAtricleBySource")
	}

	var r0 *domain.Article
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*domain.Article, error)); ok {
		return rf(ctx, sourceId, sourceSystem)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.Article); ok {
		r0 = rf(ctx, sourceId, sourceSystem)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Article)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, sourceId, sourceSystem)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLastNIdsModTimeBySource provides a mock function with given fields: ctx, n, source
func (_m *PollStorage) GetLastNIdsModTimeBySource(ctx context.Context, n int, source string) ([]domain.ArticleOriginMod, error) {
	ret := _m.Called(ctx, n, source)

	if len(ret) == 0 {
		panic("no return value specified for GetLastNIdsModTimeBySource")
	}

	var r0 []domain.ArticleOriginMod
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string) ([]domain.ArticleOriginMod, error)); ok {
		return rf(ctx, n, source)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int, string) []domain.ArticleOriginMod); ok {
		r0 = rf(ctx, n, source)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ArticleOriginMod)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, n, source)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReplaceArticleById provides a mock function with given fields: ctx, id, article
func (_m *PollStorage) ReplaceArticleById(ctx context.Context, id string, article domain.Article) error {
	ret := _m.Called(ctx, id, article)

	if len(ret) == 0 {
		panic("no return value specified for ReplaceArticleById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, domain.Article) error); ok {
		r0 = rf(ctx, id, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveArticle provides a mock function with given fields: ctx, article
func (_m *PollStorage) SaveArticle(ctx context.Context, article domain.Article) error {
	ret := _m.Called(ctx, article)

	if len(ret) == 0 {
		panic("no return value specified for SaveArticle")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Article) error); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPollStorage creates a new instance of PollStorage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPollStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *PollStorage {
	mock := &PollStorage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
