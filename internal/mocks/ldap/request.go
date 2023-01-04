// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	ber "github.com/go-asn1-ber/asn1-ber"

	mock "github.com/stretchr/testify/mock"
)

// request is an autogenerated mock type for the request type
type request struct {
	mock.Mock
}

type request_Expecter struct {
	mock *mock.Mock
}

func (_m *request) EXPECT() *request_Expecter {
	return &request_Expecter{mock: &_m.Mock}
}

// appendTo provides a mock function with given fields: _a0
func (_m *request) appendTo(_a0 *ber.Packet) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ber.Packet) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// request_appendTo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'appendTo'
type request_appendTo_Call struct {
	*mock.Call
}

// appendTo is a helper method to define mock.On call
//   - _a0 *ber.Packet
func (_e *request_Expecter) appendTo(_a0 interface{}) *request_appendTo_Call {
	return &request_appendTo_Call{Call: _e.mock.On("appendTo", _a0)}
}

func (_c *request_appendTo_Call) Run(run func(_a0 *ber.Packet)) *request_appendTo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*ber.Packet))
	})
	return _c
}

func (_c *request_appendTo_Call) Return(_a0 error) *request_appendTo_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTnewRequest interface {
	mock.TestingT
	Cleanup(func())
}

// newRequest creates a new instance of request. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newRequest(t mockConstructorTestingTnewRequest) *request {
	mock := &request{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
