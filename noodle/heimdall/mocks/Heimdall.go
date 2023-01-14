// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

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

// FindApps provides a mock function with given fields: search
func (_m *Heimdall) FindApps(search string) ([]models.ApplicationTemplate, error) {
	ret := _m.Called(search)

	var r0 []models.ApplicationTemplate
	if rf, ok := ret.Get(0).(func(string) []models.ApplicationTemplate); ok {
		r0 = rf(search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.ApplicationTemplate)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Heimdall_FindApps_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindApps'
type Heimdall_FindApps_Call struct {
	*mock.Call
}

// FindApps is a helper method to define mock.On call
//   - search string
func (_e *Heimdall_Expecter) FindApps(search interface{}) *Heimdall_FindApps_Call {
	return &Heimdall_FindApps_Call{Call: _e.mock.On("FindApps", search)}
}

func (_c *Heimdall_FindApps_Call) Run(run func(search string)) *Heimdall_FindApps_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Heimdall_FindApps_Call) Return(_a0 []models.ApplicationTemplate, _a1 error) *Heimdall_FindApps_Call {
	_c.Call.Return(_a0, _a1)
	return _c
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
