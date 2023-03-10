// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

// UserApplicationsTable is an autogenerated mock type for the UserApplicationsTable type
type UserApplicationsTable struct {
	mock.Mock
}

type UserApplicationsTable_Expecter struct {
	mock *mock.Mock
}

func (_m *UserApplicationsTable) EXPECT() *UserApplicationsTable_Expecter {
	return &UserApplicationsTable_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields:
func (_m *UserApplicationsTable) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserApplicationsTable_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserApplicationsTable_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *UserApplicationsTable_Expecter) Create() *UserApplicationsTable_Create_Call {
	return &UserApplicationsTable_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *UserApplicationsTable_Create_Call) Run(run func()) *UserApplicationsTable_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserApplicationsTable_Create_Call) Return(_a0 error) *UserApplicationsTable_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *UserApplicationsTable) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserApplicationsTable_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UserApplicationsTable_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int64
func (_e *UserApplicationsTable_Expecter) Delete(id interface{}) *UserApplicationsTable_Delete_Call {
	return &UserApplicationsTable_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *UserApplicationsTable_Delete_Call) Run(run func(id int64)) *UserApplicationsTable_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *UserApplicationsTable_Delete_Call) Return(_a0 error) *UserApplicationsTable_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *UserApplicationsTable) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserApplicationsTable_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type UserApplicationsTable_Drop_Call struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *UserApplicationsTable_Expecter) Drop() *UserApplicationsTable_Drop_Call {
	return &UserApplicationsTable_Drop_Call{Call: _e.mock.On("Drop")}
}

func (_c *UserApplicationsTable_Drop_Call) Run(run func()) *UserApplicationsTable_Drop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserApplicationsTable_Drop_Call) Return(_a0 error) *UserApplicationsTable_Drop_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetUserAllowdApplications provides a mock function with given fields: userid
func (_m *UserApplicationsTable) GetUserAllowdApplications(userid int64) (models.UsersApplications, error) {
	ret := _m.Called(userid)

	var r0 models.UsersApplications
	if rf, ok := ret.Get(0).(func(int64) models.UsersApplications); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(models.UsersApplications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserApplicationsTable_GetUserAllowdApplications_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserAllowdApplications'
type UserApplicationsTable_GetUserAllowdApplications_Call struct {
	*mock.Call
}

// GetUserAllowdApplications is a helper method to define mock.On call
//   - userid int64
func (_e *UserApplicationsTable_Expecter) GetUserAllowdApplications(userid interface{}) *UserApplicationsTable_GetUserAllowdApplications_Call {
	return &UserApplicationsTable_GetUserAllowdApplications_Call{Call: _e.mock.On("GetUserAllowdApplications", userid)}
}

func (_c *UserApplicationsTable_GetUserAllowdApplications_Call) Run(run func(userid int64)) *UserApplicationsTable_GetUserAllowdApplications_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *UserApplicationsTable_GetUserAllowdApplications_Call) Return(_a0 models.UsersApplications, _a1 error) *UserApplicationsTable_GetUserAllowdApplications_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUserApps provides a mock function with given fields: userid
func (_m *UserApplicationsTable) GetUserApps(userid int64) ([]*models.UserApplications, error) {
	ret := _m.Called(userid)

	var r0 []*models.UserApplications
	if rf, ok := ret.Get(0).(func(int64) []*models.UserApplications); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserApplications)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserApplicationsTable_GetUserApps_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserApps'
type UserApplicationsTable_GetUserApps_Call struct {
	*mock.Call
}

// GetUserApps is a helper method to define mock.On call
//   - userid int64
func (_e *UserApplicationsTable_Expecter) GetUserApps(userid interface{}) *UserApplicationsTable_GetUserApps_Call {
	return &UserApplicationsTable_GetUserApps_Call{Call: _e.mock.On("GetUserApps", userid)}
}

func (_c *UserApplicationsTable_GetUserApps_Call) Run(run func(userid int64)) *UserApplicationsTable_GetUserApps_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *UserApplicationsTable_GetUserApps_Call) Return(_a0 []*models.UserApplications, _a1 error) *UserApplicationsTable_GetUserApps_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: app
func (_m *UserApplicationsTable) Insert(app *models.UserApplications) error {
	ret := _m.Called(app)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.UserApplications) error); ok {
		r0 = rf(app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserApplicationsTable_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type UserApplicationsTable_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - app *models.UserApplications
func (_e *UserApplicationsTable_Expecter) Insert(app interface{}) *UserApplicationsTable_Insert_Call {
	return &UserApplicationsTable_Insert_Call{Call: _e.mock.On("Insert", app)}
}

func (_c *UserApplicationsTable_Insert_Call) Run(run func(app *models.UserApplications)) *UserApplicationsTable_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.UserApplications))
	})
	return _c
}

func (_c *UserApplicationsTable_Insert_Call) Return(_a0 error) *UserApplicationsTable_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upgrade provides a mock function with given fields: old_version, new_verison
func (_m *UserApplicationsTable) Upgrade(old_version int, new_verison int) error {
	ret := _m.Called(old_version, new_verison)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(old_version, new_verison)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserApplicationsTable_Upgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upgrade'
type UserApplicationsTable_Upgrade_Call struct {
	*mock.Call
}

// Upgrade is a helper method to define mock.On call
//   - old_version int
//   - new_verison int
func (_e *UserApplicationsTable_Expecter) Upgrade(old_version interface{}, new_verison interface{}) *UserApplicationsTable_Upgrade_Call {
	return &UserApplicationsTable_Upgrade_Call{Call: _e.mock.On("Upgrade", old_version, new_verison)}
}

func (_c *UserApplicationsTable_Upgrade_Call) Run(run func(old_version int, new_verison int)) *UserApplicationsTable_Upgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *UserApplicationsTable_Upgrade_Call) Return(_a0 error) *UserApplicationsTable_Upgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewUserApplicationsTable interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserApplicationsTable creates a new instance of UserApplicationsTable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserApplicationsTable(t mockConstructorTestingTNewUserApplicationsTable) *UserApplicationsTable {
	mock := &UserApplicationsTable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
