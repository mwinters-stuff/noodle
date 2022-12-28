// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	pgtype "github.com/jackc/pgx/v5/pgtype"
	mock "github.com/stretchr/testify/mock"
)

// BoolScanner is an autogenerated mock type for the BoolScanner type
type BoolScanner struct {
	mock.Mock
}

type BoolScanner_Expecter struct {
	mock *mock.Mock
}

func (_m *BoolScanner) EXPECT() *BoolScanner_Expecter {
	return &BoolScanner_Expecter{mock: &_m.Mock}
}

// ScanBool provides a mock function with given fields: v
func (_m *BoolScanner) ScanBool(v pgtype.Bool) error {
	ret := _m.Called(v)

	var r0 error
	if rf, ok := ret.Get(0).(func(pgtype.Bool) error); ok {
		r0 = rf(v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BoolScanner_ScanBool_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ScanBool'
type BoolScanner_ScanBool_Call struct {
	*mock.Call
}

// ScanBool is a helper method to define mock.On call
//   - v pgtype.Bool
func (_e *BoolScanner_Expecter) ScanBool(v interface{}) *BoolScanner_ScanBool_Call {
	return &BoolScanner_ScanBool_Call{Call: _e.mock.On("ScanBool", v)}
}

func (_c *BoolScanner_ScanBool_Call) Run(run func(v pgtype.Bool)) *BoolScanner_ScanBool_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(pgtype.Bool))
	})
	return _c
}

func (_c *BoolScanner_ScanBool_Call) Return(_a0 error) *BoolScanner_ScanBool_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewBoolScanner interface {
	mock.TestingT
	Cleanup(func())
}

// NewBoolScanner creates a new instance of BoolScanner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBoolScanner(t mockConstructorTestingTNewBoolScanner) *BoolScanner {
	mock := &BoolScanner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
