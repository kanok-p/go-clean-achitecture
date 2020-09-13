// Code generated by mockery v2.0.4. DO NOT EDIT.

package mocks

import (
	context "context"

	domainusers "github.com/kanok-p/go-clean-architecture/domain/users"
	mock "github.com/stretchr/testify/mock"

	request "github.com/kanok-p/go-clean-architecture/domain/request"

	users "github.com/kanok-p/go-clean-architecture/service/users"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *Service) Create(ctx context.Context, input *users.CreateUsers) error {
	ret := _m.Called(ctx, input)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *users.CreateUsers) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, input
func (_m *Service) Delete(ctx context.Context, input string) (*domainusers.Users, error) {
	ret := _m.Called(ctx, input)

	var r0 *domainusers.Users
	if rf, ok := ret.Get(0).(func(context.Context, string) *domainusers.Users); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainusers.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, input
func (_m *Service) Get(ctx context.Context, input string) (*domainusers.Users, error) {
	ret := _m.Called(ctx, input)

	var r0 *domainusers.Users
	if rf, ok := ret.Get(0).(func(context.Context, string) *domainusers.Users); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainusers.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, input
func (_m *Service) List(ctx context.Context, input *request.GetListInput) (int64, []*domainusers.Users, error) {
	ret := _m.Called(ctx, input)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *request.GetListInput) int64); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 []*domainusers.Users
	if rf, ok := ret.Get(1).(func(context.Context, *request.GetListInput) []*domainusers.Users); ok {
		r1 = rf(ctx, input)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*domainusers.Users)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *request.GetListInput) error); ok {
		r2 = rf(ctx, input)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Update provides a mock function with given fields: ctx, input
func (_m *Service) Update(ctx context.Context, input *users.UpdateUsers) (*domainusers.Users, error) {
	ret := _m.Called(ctx, input)

	var r0 *domainusers.Users
	if rf, ok := ret.Get(0).(func(context.Context, *users.UpdateUsers) *domainusers.Users); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domainusers.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *users.UpdateUsers) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}