// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

// UserGroupsTable is an autogenerated mock type for the UserGroupsTable type
type UserGroupsTable struct {
	mock.Mock
}

type UserGroupsTable_Expecter struct {
	mock *mock.Mock
}

func (_m *UserGroupsTable) EXPECT() *UserGroupsTable_Expecter {
	return &UserGroupsTable_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields:
func (_m *UserGroupsTable) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGroupsTable_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type UserGroupsTable_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *UserGroupsTable_Expecter) Create() *UserGroupsTable_Create_Call {
	return &UserGroupsTable_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *UserGroupsTable_Create_Call) Run(run func()) *UserGroupsTable_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserGroupsTable_Create_Call) Return(_a0 error) *UserGroupsTable_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: user
func (_m *UserGroupsTable) Delete(user models.UserGroup) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.UserGroup) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGroupsTable_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type UserGroupsTable_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - user models.UserGroup
func (_e *UserGroupsTable_Expecter) Delete(user interface{}) *UserGroupsTable_Delete_Call {
	return &UserGroupsTable_Delete_Call{Call: _e.mock.On("Delete", user)}
}

func (_c *UserGroupsTable_Delete_Call) Run(run func(user models.UserGroup)) *UserGroupsTable_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.UserGroup))
	})
	return _c
}

func (_c *UserGroupsTable_Delete_Call) Return(_a0 error) *UserGroupsTable_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *UserGroupsTable) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGroupsTable_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type UserGroupsTable_Drop_Call struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *UserGroupsTable_Expecter) Drop() *UserGroupsTable_Drop_Call {
	return &UserGroupsTable_Drop_Call{Call: _e.mock.On("Drop")}
}

func (_c *UserGroupsTable_Drop_Call) Run(run func()) *UserGroupsTable_Drop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserGroupsTable_Drop_Call) Return(_a0 error) *UserGroupsTable_Drop_Call {
	_c.Call.Return(_a0)
	return _c
}

// Exists provides a mock function with given fields: groupid, userid
func (_m *UserGroupsTable) Exists(groupid int64, userid int64) (bool, error) {
	ret := _m.Called(groupid, userid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, int64) bool); ok {
		r0 = rf(groupid, userid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64, int64) error); ok {
		r1 = rf(groupid, userid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGroupsTable_Exists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exists'
type UserGroupsTable_Exists_Call struct {
	*mock.Call
}

// Exists is a helper method to define mock.On call
//   - groupid int64
//   - userid int64
func (_e *UserGroupsTable_Expecter) Exists(groupid interface{}, userid interface{}) *UserGroupsTable_Exists_Call {
	return &UserGroupsTable_Exists_Call{Call: _e.mock.On("Exists", groupid, userid)}
}

func (_c *UserGroupsTable_Exists_Call) Run(run func(groupid int64, userid int64)) *UserGroupsTable_Exists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *UserGroupsTable_Exists_Call) Return(_a0 bool, _a1 error) *UserGroupsTable_Exists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *UserGroupsTable) GetAll() ([]*models.UserGroup, error) {
	ret := _m.Called()

	var r0 []*models.UserGroup
	if rf, ok := ret.Get(0).(func() []*models.UserGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGroupsTable_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type UserGroupsTable_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *UserGroupsTable_Expecter) GetAll() *UserGroupsTable_GetAll_Call {
	return &UserGroupsTable_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *UserGroupsTable_GetAll_Call) Run(run func()) *UserGroupsTable_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *UserGroupsTable_GetAll_Call) Return(_a0 []*models.UserGroup, _a1 error) *UserGroupsTable_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetGroup provides a mock function with given fields: groupid
func (_m *UserGroupsTable) GetGroup(groupid int64) ([]*models.UserGroup, error) {
	ret := _m.Called(groupid)

	var r0 []*models.UserGroup
	if rf, ok := ret.Get(0).(func(int64) []*models.UserGroup); ok {
		r0 = rf(groupid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(groupid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserGroupsTable_GetGroup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroup'
type UserGroupsTable_GetGroup_Call struct {
	*mock.Call
}

// GetGroup is a helper method to define mock.On call
//   - groupid int64
func (_e *UserGroupsTable_Expecter) GetGroup(groupid interface{}) *UserGroupsTable_GetGroup_Call {
	return &UserGroupsTable_GetGroup_Call{Call: _e.mock.On("GetGroup", groupid)}
}

func (_c *UserGroupsTable_GetGroup_Call) Run(run func(groupid int64)) *UserGroupsTable_GetGroup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *UserGroupsTable_GetGroup_Call) Return(_a0 []*models.UserGroup, _a1 error) *UserGroupsTable_GetGroup_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUser provides a mock function with given fields: userid
func (_m *UserGroupsTable) GetUser(userid int64) ([]*models.UserGroup, error) {
	ret := _m.Called(userid)

	var r0 []*models.UserGroup
	if rf, ok := ret.Get(0).(func(int64) []*models.UserGroup); ok {
		r0 = rf(userid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.UserGroup)
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

// UserGroupsTable_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type UserGroupsTable_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - userid int64
func (_e *UserGroupsTable_Expecter) GetUser(userid interface{}) *UserGroupsTable_GetUser_Call {
	return &UserGroupsTable_GetUser_Call{Call: _e.mock.On("GetUser", userid)}
}

func (_c *UserGroupsTable_GetUser_Call) Run(run func(userid int64)) *UserGroupsTable_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *UserGroupsTable_GetUser_Call) Return(_a0 []*models.UserGroup, _a1 error) *UserGroupsTable_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: user
func (_m *UserGroupsTable) Insert(user *models.UserGroup) error {
	ret := _m.Called(user)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.UserGroup) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGroupsTable_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type UserGroupsTable_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - user *models.UserGroup
func (_e *UserGroupsTable_Expecter) Insert(user interface{}) *UserGroupsTable_Insert_Call {
	return &UserGroupsTable_Insert_Call{Call: _e.mock.On("Insert", user)}
}

func (_c *UserGroupsTable_Insert_Call) Run(run func(user *models.UserGroup)) *UserGroupsTable_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.UserGroup))
	})
	return _c
}

func (_c *UserGroupsTable_Insert_Call) Return(_a0 error) *UserGroupsTable_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upgrade provides a mock function with given fields: old_version, new_verison
func (_m *UserGroupsTable) Upgrade(old_version int, new_verison int) error {
	ret := _m.Called(old_version, new_verison)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(old_version, new_verison)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserGroupsTable_Upgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upgrade'
type UserGroupsTable_Upgrade_Call struct {
	*mock.Call
}

// Upgrade is a helper method to define mock.On call
//   - old_version int
//   - new_verison int
func (_e *UserGroupsTable_Expecter) Upgrade(old_version interface{}, new_verison interface{}) *UserGroupsTable_Upgrade_Call {
	return &UserGroupsTable_Upgrade_Call{Call: _e.mock.On("Upgrade", old_version, new_verison)}
}

func (_c *UserGroupsTable_Upgrade_Call) Run(run func(old_version int, new_verison int)) *UserGroupsTable_Upgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *UserGroupsTable_Upgrade_Call) Return(_a0 error) *UserGroupsTable_Upgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewUserGroupsTable interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserGroupsTable creates a new instance of UserGroupsTable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserGroupsTable(t mockConstructorTestingTNewUserGroupsTable) *UserGroupsTable {
	mock := &UserGroupsTable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
