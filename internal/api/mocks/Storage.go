// Code generated by mockery v2.24.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "gitlab.com/pet-pr-social-network/relation-service/internal/model"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// CreateCity provides a mock function with given fields: ctx, city
func (_m *Storage) CreateCity(ctx context.Context, city model.City) (int64, error) {
	ret := _m.Called(ctx, city)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.City) (int64, error)); ok {
		return rf(ctx, city)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.City) int64); ok {
		r0 = rf(ctx, city)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.City) error); ok {
		r1 = rf(ctx, city)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateInterest provides a mock function with given fields: ctx, interest
func (_m *Storage) CreateInterest(ctx context.Context, interest model.Interest) (int64, error) {
	ret := _m.Called(ctx, interest)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.Interest) (int64, error)); ok {
		return rf(ctx, interest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.Interest) int64); ok {
		r0 = rf(ctx, interest)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.Interest) error); ok {
		r1 = rf(ctx, interest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *Storage) CreateUser(ctx context.Context, user model.User) (int64, error) {
	ret := _m.Called(ctx, user)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.User) (int64, error)); ok {
		return rf(ctx, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.User) int64); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCities provides a mock function with given fields: ctx
func (_m *Storage) GetAllCities(ctx context.Context) ([]model.City, error) {
	ret := _m.Called(ctx)

	var r0 []model.City
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.City, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.City); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.City)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllInterests provides a mock function with given fields: ctx
func (_m *Storage) GetAllInterests(ctx context.Context) ([]model.Interest, error) {
	ret := _m.Called(ctx)

	var r0 []model.Interest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]model.Interest, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []model.Interest); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Interest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCity provides a mock function with given fields: ctx, id
func (_m *Storage) GetCity(ctx context.Context, id int64) (model.City, error) {
	ret := _m.Called(ctx, id)

	var r0 model.City
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.City, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.City); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.City)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInterest provides a mock function with given fields: ctx, id
func (_m *Storage) GetInterest(ctx context.Context, id int64) (model.Interest, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Interest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.Interest, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.Interest); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Interest)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id
func (_m *Storage) GetUser(ctx context.Context, id int64) (model.User, error) {
	ret := _m.Called(ctx, id)

	var r0 model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (model.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) model.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
