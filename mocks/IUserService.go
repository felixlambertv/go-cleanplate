// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	model "github.com/felixlambertv/go-cleanplate/internal/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	request "github.com/felixlambertv/go-cleanplate/internal/controller/request"

	service "github.com/felixlambertv/go-cleanplate/internal/service"
)

// IUserService is an autogenerated mock type for the IUserService type
type IUserService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: req
func (_m *IUserService) CreateUser(req request.CreateUserRequest) (*model.User, error) {
	ret := _m.Called(req)

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(request.CreateUserRequest) (*model.User, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(request.CreateUserRequest) *model.User); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(request.CreateUserRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields:
func (_m *IUserService) GetUsers() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]model.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTrx provides a mock function with given fields: trxHandle
func (_m *IUserService) WithTrx(trxHandle *gorm.DB) service.IUserService {
	ret := _m.Called(trxHandle)

	var r0 service.IUserService
	if rf, ok := ret.Get(0).(func(*gorm.DB) service.IUserService); ok {
		r0 = rf(trxHandle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(service.IUserService)
		}
	}

	return r0
}

type mockConstructorTestingTNewIUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserService creates a new instance of IUserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserService(t mockConstructorTestingTNewIUserService) *IUserService {
	mock := &IUserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
