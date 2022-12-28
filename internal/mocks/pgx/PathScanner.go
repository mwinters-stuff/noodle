// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	pgtype "github.com/jackc/pgx/v5/pgtype"
	mock "github.com/stretchr/testify/mock"
)

// PathScanner is an autogenerated mock type for the PathScanner type
type PathScanner struct {
	mock.Mock
}

type PathScanner_Expecter struct {
	mock *mock.Mock
}

func (_m *PathScanner) EXPECT() *PathScanner_Expecter {
	return &PathScanner_Expecter{mock: &_m.Mock}
}

// ScanPath provides a mock function with given fields: v
func (_m *PathScanner) ScanPath(v pgtype.Path) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(pgtype.Path) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// PathScanner_ScanPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScanPath'
type PathScanner_ScanPath_Call struct {
	*mock.Call
}

// ScanPath is a helper method to define mock.On call
//   - v pgtype.Path
func (_e *PathScanner_Expecter) ScanPath(v interface{}) *PathScanner_ScanPath_Call {
	return &PathScanner_ScanPath_Call{Call: _e.mock.On("ScanPath", v)}
}

func (_c *PathScanner_ScanPath_Call) Run(run func(v pgtype.Path)) *PathScanner_ScanPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pgtype.Path))
	})
	return _c
}

func (_c *PathScanner_ScanPath_Call) Return(_a0 error) *PathScanner_ScanPath_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewPathScanner interface {
	mock.TestingT
	Cleanup(func())
}

// NewPathScanner creates a new instance of PathScanner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPathScanner(t mockConstructorTestingTNewPathScanner) *PathScanner {
	mock := &PathScanner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
