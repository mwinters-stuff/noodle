// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Heimdall is an autogenerated mock type for the Heimdall type
type Heimdall struct {
	mock.Mock
}

type Heimdall_Expecter struct {
	mock *mock.Mock
}

func (_m *Heimdall) EXPECT() *Heimdall_Expecter {
	return &Heimdall_Expecter{mock: &_m.Mock}
}

// UpdateFromServer provides a mock function with given fields:
func (_m *Heimdall) UpdateFromServer() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Heimdall_UpdateFromServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateFromServer'
type Heimdall_UpdateFromServer_Call struct {
	*mock.Call
}

// UpdateFromServer is a helper method to define mock.On call
func (_e *Heimdall_Expecter) UpdateFromServer() *Heimdall_UpdateFromServer_Call {
	return &Heimdall_UpdateFromServer_Call{Call: _e.mock.On("UpdateFromServer")}
}

func (_c *Heimdall_UpdateFromServer_Call) Run(run func()) *Heimdall_UpdateFromServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Heimdall_UpdateFromServer_Call) Return(_a0 error) *Heimdall_UpdateFromServer_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewHeimdall interface {
	mock.TestingT
	Cleanup(func())
}

// NewHeimdall creates a new instance of Heimdall. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHeimdall(t mockConstructorTestingTNewHeimdall) *Heimdall {
	mock := &Heimdall{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
