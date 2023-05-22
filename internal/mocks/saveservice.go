// Code generated by mockery v2.27.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	domain "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

// SaveService is an autogenerated mock type for the SaveService type
type SaveService struct {
	mock.Mock
}

// Execute provides a mock function with given fields: title, description, price
func (_m *SaveService) Execute(title string, description string, price float64) domain.Ad {
	ret := _m.Called(title, description, price)

	var r0 domain.Ad
	if rf, ok := ret.Get(0).(func(string, string, float64) domain.Ad); ok {
		r0 = rf(title, description, price)
	} else {
		r0 = ret.Get(0).(domain.Ad)
	}

	return r0
}

type mockConstructorTestingTNewSaveService interface {
	mock.TestingT
	Cleanup(func())
}

// NewSaveService creates a new instance of SaveService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSaveService(t mockConstructorTestingTNewSaveService) *SaveService {
	mock := &SaveService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}