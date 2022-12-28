// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FrontendMessage is an autogenerated mock type for the FrontendMessage type
type FrontendMessage struct {
	mock.Mock
}

type FrontendMessage_Expecter struct {
	mock *mock.Mock
}

func (_m *FrontendMessage) EXPECT() *FrontendMessage_Expecter {
	return &FrontendMessage_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: data
func (_m *FrontendMessage) Decode(data []byte) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FrontendMessage_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type FrontendMessage_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - data []byte
func (_e *FrontendMessage_Expecter) Decode(data interface{}) *FrontendMessage_Decode_Call {
	return &FrontendMessage_Decode_Call{Call: _e.mock.On("Decode", data)}
}

func (_c *FrontendMessage_Decode_Call) Run(run func(data []byte)) *FrontendMessage_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *FrontendMessage_Decode_Call) Return(_a0 error) *FrontendMessage_Decode_Call {
	_c.Call.Return(_a0)
	return _c
}

// Encode provides a mock function with given fields: dst
func (_m *FrontendMessage) Encode(dst []byte) []byte {
	ret := _m.Called(dst)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(dst)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// FrontendMessage_Encode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Encode'
type FrontendMessage_Encode_Call struct {
	*mock.Call
}

// Encode is a helper method to define mock.On call
//   - dst []byte
func (_e *FrontendMessage_Expecter) Encode(dst interface{}) *FrontendMessage_Encode_Call {
	return &FrontendMessage_Encode_Call{Call: _e.mock.On("Encode", dst)}
}

func (_c *FrontendMessage_Encode_Call) Run(run func(dst []byte)) *FrontendMessage_Encode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *FrontendMessage_Encode_Call) Return(_a0 []byte) *FrontendMessage_Encode_Call {
	_c.Call.Return(_a0)
	return _c
}

// Frontend provides a mock function with given fields:
func (_m *FrontendMessage) Frontend() {
	_m.Called()
}

// FrontendMessage_Frontend_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Frontend'
type FrontendMessage_Frontend_Call struct {
	*mock.Call
}

// Frontend is a helper method to define mock.On call
func (_e *FrontendMessage_Expecter) Frontend() *FrontendMessage_Frontend_Call {
	return &FrontendMessage_Frontend_Call{Call: _e.mock.On("Frontend")}
}

func (_c *FrontendMessage_Frontend_Call) Run(run func()) *FrontendMessage_Frontend_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FrontendMessage_Frontend_Call) Return() *FrontendMessage_Frontend_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewFrontendMessage interface {
	mock.TestingT
	Cleanup(func())
}

// NewFrontendMessage creates a new instance of FrontendMessage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewFrontendMessage(t mockConstructorTestingTNewFrontendMessage) *FrontendMessage {
	mock := &FrontendMessage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
