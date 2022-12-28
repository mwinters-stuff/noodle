// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// OptionOpenDB is an autogenerated mock type for the OptionOpenDB type
type OptionOpenDB struct {
	mock.Mock
}

type OptionOpenDB_Expecter struct {
	mock *mock.Mock
}

func (_m *OptionOpenDB) EXPECT() *OptionOpenDB_Expecter {
	return &OptionOpenDB_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0
func (_m *OptionOpenDB) Execute(_a0 *stdlib.connector) {
	_m.Called(_a0)
}

// OptionOpenDB_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type OptionOpenDB_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 *stdlib.connector
func (_e *OptionOpenDB_Expecter) Execute(_a0 interface{}) *OptionOpenDB_Execute_Call {
	return &OptionOpenDB_Execute_Call{Call: _e.mock.On("Execute", _a0)}
}

func (_c *OptionOpenDB_Execute_Call) Run(run func(_a0 *stdlib.connector)) *OptionOpenDB_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*stdlib.connector))
	})
	return _c
}

func (_c *OptionOpenDB_Execute_Call) Return() *OptionOpenDB_Execute_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTNewOptionOpenDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewOptionOpenDB creates a new instance of OptionOpenDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOptionOpenDB(t mockConstructorTestingTNewOptionOpenDB) *OptionOpenDB {
	mock := &OptionOpenDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
