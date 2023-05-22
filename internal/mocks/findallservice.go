// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

// FindAdsService is an autogenerated mock type for the FindAdsService type
type FindAdsService struct {
	mock.Mock
}

// Execute provides a mock function with given fields:
func (_m *FindAdsService) Execute() []domain.Ad {
	ret := _m.Called()

	var r0 []domain.Ad
	if rf, ok := ret.Get(0).(func() []domain.Ad); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Ad)
		}
	}

	return r0
}

type mockConstructorTestingTNewFindAdsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewFindAdsService creates a new instance of FindAdsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFindAdsService(t mockConstructorTestingTNewFindAdsService) *FindAdsService {
	mock := &FindAdsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
