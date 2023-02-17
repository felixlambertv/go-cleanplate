// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	model "github.com/felixlambertv/go-cleanplate/internal/model"
	mock "github.com/stretchr/testify/mock"
	gorm "gorm.io/gorm"

	repository "github.com/felixlambertv/go-cleanplate/internal/repository"
)

// IUserRepo is an autogenerated mock type for the IUserRepo type
type IUserRepo struct {
	mock.Mock
}

// FindAll provides a mock function with given fields:
func (_m *IUserRepo) FindAll() ([]model.User, error) {
	ret := _m.Called()

	var r0 []model.User
	if rf, ok := ret.Get(0).(func() []model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields:
func (_m *IUserRepo) Store() (*model.User, error) {
	ret := _m.Called()

	var r0 *model.User
	if rf, ok := ret.Get(0).(func() *model.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WithTrx provides a mock function with given fields: trxHandle
func (_m *IUserRepo) WithTrx(trxHandle *gorm.DB) repository.IUserRepo {
	ret := _m.Called(trxHandle)

	var r0 repository.IUserRepo
	if rf, ok := ret.Get(0).(func(*gorm.DB) repository.IUserRepo); ok {
		r0 = rf(trxHandle)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.IUserRepo)
		}
	}

	return r0
}

type mockConstructorTestingTNewIUserRepo interface {
	mock.TestingT
	Cleanup(func())
}

// NewIUserRepo creates a new instance of IUserRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIUserRepo(t mockConstructorTestingTNewIUserRepo) *IUserRepo {
	mock := &IUserRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
