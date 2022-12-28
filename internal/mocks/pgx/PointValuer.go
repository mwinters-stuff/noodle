// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	pgtype "github.com/jackc/pgx/v5/pgtype"
	mock "github.com/stretchr/testify/mock"
)

// PointValuer is an autogenerated mock type for the PointValuer type
type PointValuer struct {
	mock.Mock
}

type PointValuer_Expecter struct {
	mock *mock.Mock
}

func (_m *PointValuer) EXPECT() *PointValuer_Expecter {
	return &PointValuer_Expecter{mock: &_m.Mock}
}

// PointValue provides a mock function with given fields:
func (_m *PointValuer) PointValue() (pgtype.Point, error) {
	ret := _m.Called()

	var r0 pgtype.Point
	if rf, ok := ret.Get(0).(func() pgtype.Point); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(pgtype.Point)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PointValuer_PointValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PointValue'
type PointValuer_PointValue_Call struct {
	*mock.Call
}

// PointValue is a helper method to define mock.On call
func (_e *PointValuer_Expecter) PointValue() *PointValuer_PointValue_Call {
	return &PointValuer_PointValue_Call{Call: _e.mock.On("PointValue")}
}

func (_c *PointValuer_PointValue_Call) Run(run func()) *PointValuer_PointValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *PointValuer_PointValue_Call) Return(_a0 pgtype.Point, _a1 error) *PointValuer_PointValue_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewPointValuer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPointValuer creates a new instance of PointValuer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPointValuer(t mockConstructorTestingTNewPointValuer) *PointValuer {
	mock := &PointValuer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
