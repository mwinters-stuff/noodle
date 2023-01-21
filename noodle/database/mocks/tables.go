// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	database "github.com/mwinters-stuff/noodle/noodle/database"
	mock "github.com/stretchr/testify/mock"
)

// Tables is an autogenerated mock type for the Tables type
type Tables struct {
	mock.Mock
}

type Tables_Expecter struct {
	mock *mock.Mock
}

func (_m *Tables) EXPECT() *Tables_Expecter {
	return &Tables_Expecter{mock: &_m.Mock}
}

// AppTemplateTable provides a mock function with given fields:
func (_m *Tables) AppTemplateTable() database.AppTemplateTable {
	ret := _m.Called()

	var r0 database.AppTemplateTable
	if rf, ok := ret.Get(0).(func() database.AppTemplateTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.AppTemplateTable)
		}
	}

	return r0
}

// Tables_AppTemplateTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AppTemplateTable'
type Tables_AppTemplateTable_Call struct {
	*mock.Call
}

// AppTemplateTable is a helper method to define mock.On call
func (_e *Tables_Expecter) AppTemplateTable() *Tables_AppTemplateTable_Call {
	return &Tables_AppTemplateTable_Call{Call: _e.mock.On("AppTemplateTable")}
}

func (_c *Tables_AppTemplateTable_Call) Run(run func()) *Tables_AppTemplateTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_AppTemplateTable_Call) Return(_a0 database.AppTemplateTable) *Tables_AppTemplateTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// ApplicationTabTable provides a mock function with given fields:
func (_m *Tables) ApplicationTabTable() database.ApplicationTabTable {
	ret := _m.Called()

	var r0 database.ApplicationTabTable
	if rf, ok := ret.Get(0).(func() database.ApplicationTabTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.ApplicationTabTable)
		}
	}

	return r0
}

// Tables_ApplicationTabTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplicationTabTable'
type Tables_ApplicationTabTable_Call struct {
	*mock.Call
}

// ApplicationTabTable is a helper method to define mock.On call
func (_e *Tables_Expecter) ApplicationTabTable() *Tables_ApplicationTabTable_Call {
	return &Tables_ApplicationTabTable_Call{Call: _e.mock.On("ApplicationTabTable")}
}

func (_c *Tables_ApplicationTabTable_Call) Run(run func()) *Tables_ApplicationTabTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_ApplicationTabTable_Call) Return(_a0 database.ApplicationTabTable) *Tables_ApplicationTabTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// ApplicationsTable provides a mock function with given fields:
func (_m *Tables) ApplicationsTable() database.ApplicationsTable {
	ret := _m.Called()

	var r0 database.ApplicationsTable
	if rf, ok := ret.Get(0).(func() database.ApplicationsTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.ApplicationsTable)
		}
	}

	return r0
}

// Tables_ApplicationsTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplicationsTable'
type Tables_ApplicationsTable_Call struct {
	*mock.Call
}

// ApplicationsTable is a helper method to define mock.On call
func (_e *Tables_Expecter) ApplicationsTable() *Tables_ApplicationsTable_Call {
	return &Tables_ApplicationsTable_Call{Call: _e.mock.On("ApplicationsTable")}
}

func (_c *Tables_ApplicationsTable_Call) Run(run func()) *Tables_ApplicationsTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_ApplicationsTable_Call) Return(_a0 database.ApplicationsTable) *Tables_ApplicationsTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// Create provides a mock function with given fields:
func (_m *Tables) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tables_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type Tables_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *Tables_Expecter) Create() *Tables_Create_Call {
	return &Tables_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *Tables_Create_Call) Run(run func()) *Tables_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_Create_Call) Return(_a0 error) *Tables_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *Tables) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Tables_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type Tables_Drop_Call struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *Tables_Expecter) Drop() *Tables_Drop_Call {
	return &Tables_Drop_Call{Call: _e.mock.On("Drop")}
}

func (_c *Tables_Drop_Call) Run(run func()) *Tables_Drop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_Drop_Call) Return(_a0 error) *Tables_Drop_Call {
	_c.Call.Return(_a0)
	return _c
}

// GroupApplicationsTable provides a mock function with given fields:
func (_m *Tables) GroupApplicationsTable() database.GroupApplicationsTable {
	ret := _m.Called()

	var r0 database.GroupApplicationsTable
	if rf, ok := ret.Get(0).(func() database.GroupApplicationsTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.GroupApplicationsTable)
		}
	}

	return r0
}

// Tables_GroupApplicationsTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GroupApplicationsTable'
type Tables_GroupApplicationsTable_Call struct {
	*mock.Call
}

// GroupApplicationsTable is a helper method to define mock.On call
func (_e *Tables_Expecter) GroupApplicationsTable() *Tables_GroupApplicationsTable_Call {
	return &Tables_GroupApplicationsTable_Call{Call: _e.mock.On("GroupApplicationsTable")}
}

func (_c *Tables_GroupApplicationsTable_Call) Run(run func()) *Tables_GroupApplicationsTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_GroupApplicationsTable_Call) Return(_a0 database.GroupApplicationsTable) *Tables_GroupApplicationsTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// GroupTable provides a mock function with given fields:
func (_m *Tables) GroupTable() database.GroupTable {
	ret := _m.Called()

	var r0 database.GroupTable
	if rf, ok := ret.Get(0).(func() database.GroupTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.GroupTable)
		}
	}

	return r0
}

// Tables_GroupTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GroupTable'
type Tables_GroupTable_Call struct {
	*mock.Call
}

// GroupTable is a helper method to define mock.On call
func (_e *Tables_Expecter) GroupTable() *Tables_GroupTable_Call {
	return &Tables_GroupTable_Call{Call: _e.mock.On("GroupTable")}
}

func (_c *Tables_GroupTable_Call) Run(run func()) *Tables_GroupTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_GroupTable_Call) Return(_a0 database.GroupTable) *Tables_GroupTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// InitTables provides a mock function with given fields: db
func (_m *Tables) InitTables(db database.Database) {
	_m.Called(db)
}

// Tables_InitTables_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitTables'
type Tables_InitTables_Call struct {
	*mock.Call
}

// InitTables is a helper method to define mock.On call
//   - db database.Database
func (_e *Tables_Expecter) InitTables(db interface{}) *Tables_InitTables_Call {
	return &Tables_InitTables_Call{Call: _e.mock.On("InitTables", db)}
}

func (_c *Tables_InitTables_Call) Run(run func(db database.Database)) *Tables_InitTables_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(database.Database))
	})
	return _c
}

func (_c *Tables_InitTables_Call) Return() *Tables_InitTables_Call {
	_c.Call.Return()
	return _c
}

// TabTable provides a mock function with given fields:
func (_m *Tables) TabTable() database.TabTable {
	ret := _m.Called()

	var r0 database.TabTable
	if rf, ok := ret.Get(0).(func() database.TabTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.TabTable)
		}
	}

	return r0
}

// Tables_TabTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'TabTable'
type Tables_TabTable_Call struct {
	*mock.Call
}

// TabTable is a helper method to define mock.On call
func (_e *Tables_Expecter) TabTable() *Tables_TabTable_Call {
	return &Tables_TabTable_Call{Call: _e.mock.On("TabTable")}
}

func (_c *Tables_TabTable_Call) Run(run func()) *Tables_TabTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_TabTable_Call) Return(_a0 database.TabTable) *Tables_TabTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// UserApplicationsTable provides a mock function with given fields:
func (_m *Tables) UserApplicationsTable() database.UserApplicationsTable {
	ret := _m.Called()

	var r0 database.UserApplicationsTable
	if rf, ok := ret.Get(0).(func() database.UserApplicationsTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UserApplicationsTable)
		}
	}

	return r0
}

// Tables_UserApplicationsTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserApplicationsTable'
type Tables_UserApplicationsTable_Call struct {
	*mock.Call
}

// UserApplicationsTable is a helper method to define mock.On call
func (_e *Tables_Expecter) UserApplicationsTable() *Tables_UserApplicationsTable_Call {
	return &Tables_UserApplicationsTable_Call{Call: _e.mock.On("UserApplicationsTable")}
}

func (_c *Tables_UserApplicationsTable_Call) Run(run func()) *Tables_UserApplicationsTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_UserApplicationsTable_Call) Return(_a0 database.UserApplicationsTable) *Tables_UserApplicationsTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// UserGroupsTable provides a mock function with given fields:
func (_m *Tables) UserGroupsTable() database.UserGroupsTable {
	ret := _m.Called()

	var r0 database.UserGroupsTable
	if rf, ok := ret.Get(0).(func() database.UserGroupsTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UserGroupsTable)
		}
	}

	return r0
}

// Tables_UserGroupsTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserGroupsTable'
type Tables_UserGroupsTable_Call struct {
	*mock.Call
}

// UserGroupsTable is a helper method to define mock.On call
func (_e *Tables_Expecter) UserGroupsTable() *Tables_UserGroupsTable_Call {
	return &Tables_UserGroupsTable_Call{Call: _e.mock.On("UserGroupsTable")}
}

func (_c *Tables_UserGroupsTable_Call) Run(run func()) *Tables_UserGroupsTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_UserGroupsTable_Call) Return(_a0 database.UserGroupsTable) *Tables_UserGroupsTable_Call {
	_c.Call.Return(_a0)
	return _c
}

// UserTable provides a mock function with given fields:
func (_m *Tables) UserTable() database.UserTable {
	ret := _m.Called()

	var r0 database.UserTable
	if rf, ok := ret.Get(0).(func() database.UserTable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.UserTable)
		}
	}

	return r0
}

// Tables_UserTable_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserTable'
type Tables_UserTable_Call struct {
	*mock.Call
}

// UserTable is a helper method to define mock.On call
func (_e *Tables_Expecter) UserTable() *Tables_UserTable_Call {
	return &Tables_UserTable_Call{Call: _e.mock.On("UserTable")}
}

func (_c *Tables_UserTable_Call) Run(run func()) *Tables_UserTable_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Tables_UserTable_Call) Return(_a0 database.UserTable) *Tables_UserTable_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewTables interface {
	mock.TestingT
	Cleanup(func())
}

// NewTables creates a new instance of Tables. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTables(t mockConstructorTestingTNewTables) *Tables {
	mock := &Tables{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}