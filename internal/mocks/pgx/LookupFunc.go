// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// LookupFunc is an autogenerated mock type for the LookupFunc type
type LookupFunc struct {
	mock.Mock
}

type LookupFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *LookupFunc) EXPECT() *LookupFunc_Expecter {
	return &LookupFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, host
func (_m *LookupFunc) Execute(ctx context.Context, host string) ([]string, error) {
	ret := _m.Called(ctx, host)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string) []string); ok {
		r0 = rf(ctx, host)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, host)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LookupFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type LookupFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - host string
func (_e *LookupFunc_Expecter) Execute(ctx interface{}, host interface{}) *LookupFunc_Execute_Call {
	return &LookupFunc_Execute_Call{Call: _e.mock.On("Execute", ctx, host)}
}

func (_c *LookupFunc_Execute_Call) Run(run func(ctx context.Context, host string)) *LookupFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *LookupFunc_Execute_Call) Return(addrs []string, err error) *LookupFunc_Execute_Call {
	_c.Call.Return(addrs, err)
	return _c
}

type mockConstructorTestingTNewLookupFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewLookupFunc creates a new instance of LookupFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLookupFunc(t mockConstructorTestingTNewLookupFunc) *LookupFunc {
	mock := &LookupFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
