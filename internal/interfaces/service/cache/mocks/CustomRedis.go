// Code generated by mockery v2.13.1. DO NOT EDIT.

package mocks

import (
	context "context"

	redis "github.com/go-redis/redis/v9"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// CustomRedis is an autogenerated mock type for the CustomRedis type
type CustomRedis struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, key
func (_m *CustomRedis) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := _m.Called(ctx, key)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func(context.Context, string) *redis.StringCmd); ok {
		r0 = rf(ctx, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

// Set provides a mock function with given fields: ctx, key, value, expiration
func (_m *CustomRedis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := _m.Called(ctx, key, value, expiration)

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}, time.Duration) *redis.StatusCmd); ok {
		r0 = rf(ctx, key, value, expiration)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}

type mockConstructorTestingTNewCustomRedis interface {
	mock.TestingT
	Cleanup(func())
}

// NewCustomRedis creates a new instance of CustomRedis. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCustomRedis(t mockConstructorTestingTNewCustomRedis) *CustomRedis {
	mock := &CustomRedis{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
