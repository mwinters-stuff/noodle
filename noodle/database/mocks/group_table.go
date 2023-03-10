// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

// GroupTable is an autogenerated mock type for the GroupTable type
type GroupTable struct {
	mock.Mock
}

type GroupTable_Expecter struct {
	mock *mock.Mock
}

func (_m *GroupTable) EXPECT() *GroupTable_Expecter {
	return &GroupTable_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields:
func (_m *GroupTable) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type GroupTable_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *GroupTable_Expecter) Create() *GroupTable_Create_Call {
	return &GroupTable_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *GroupTable_Create_Call) Run(run func()) *GroupTable_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GroupTable_Create_Call) Return(_a0 error) *GroupTable_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: group
func (_m *GroupTable) Delete(group models.Group) error {
	ret := _m.Called(group)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Group) error); ok {
		r0 = rf(group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type GroupTable_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - group models.Group
func (_e *GroupTable_Expecter) Delete(group interface{}) *GroupTable_Delete_Call {
	return &GroupTable_Delete_Call{Call: _e.mock.On("Delete", group)}
}

func (_c *GroupTable_Delete_Call) Run(run func(group models.Group)) *GroupTable_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Group))
	})
	return _c
}

func (_c *GroupTable_Delete_Call) Return(_a0 error) *GroupTable_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *GroupTable) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type GroupTable_Drop_Call struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *GroupTable_Expecter) Drop() *GroupTable_Drop_Call {
	return &GroupTable_Drop_Call{Call: _e.mock.On("Drop")}
}

func (_c *GroupTable_Drop_Call) Run(run func()) *GroupTable_Drop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GroupTable_Drop_Call) Return(_a0 error) *GroupTable_Drop_Call {
	_c.Call.Return(_a0)
	return _c
}

// ExistsDN provides a mock function with given fields: dn
func (_m *GroupTable) ExistsDN(dn string) (bool, error) {
	ret := _m.Called(dn)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(dn)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupTable_ExistsDN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExistsDN'
type GroupTable_ExistsDN_Call struct {
	*mock.Call
}

// ExistsDN is a helper method to define mock.On call
//   - dn string
func (_e *GroupTable_Expecter) ExistsDN(dn interface{}) *GroupTable_ExistsDN_Call {
	return &GroupTable_ExistsDN_Call{Call: _e.mock.On("ExistsDN", dn)}
}

func (_c *GroupTable_ExistsDN_Call) Run(run func(dn string)) *GroupTable_ExistsDN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GroupTable_ExistsDN_Call) Return(_a0 bool, _a1 error) *GroupTable_ExistsDN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ExistsName provides a mock function with given fields: groupname
func (_m *GroupTable) ExistsName(groupname string) (bool, error) {
	ret := _m.Called(groupname)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(groupname)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(groupname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupTable_ExistsName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExistsName'
type GroupTable_ExistsName_Call struct {
	*mock.Call
}

// ExistsName is a helper method to define mock.On call
//   - groupname string
func (_e *GroupTable_Expecter) ExistsName(groupname interface{}) *GroupTable_ExistsName_Call {
	return &GroupTable_ExistsName_Call{Call: _e.mock.On("ExistsName", groupname)}
}

func (_c *GroupTable_ExistsName_Call) Run(run func(groupname string)) *GroupTable_ExistsName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GroupTable_ExistsName_Call) Return(_a0 bool, _a1 error) *GroupTable_ExistsName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetAll provides a mock function with given fields:
func (_m *GroupTable) GetAll() ([]*models.Group, error) {
	ret := _m.Called()

	var r0 []*models.Group
	if rf, ok := ret.Get(0).(func() []*models.Group); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Group)
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

// GroupTable_GetAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetAll'
type GroupTable_GetAll_Call struct {
	*mock.Call
}

// GetAll is a helper method to define mock.On call
func (_e *GroupTable_Expecter) GetAll() *GroupTable_GetAll_Call {
	return &GroupTable_GetAll_Call{Call: _e.mock.On("GetAll")}
}

func (_c *GroupTable_GetAll_Call) Run(run func()) *GroupTable_GetAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *GroupTable_GetAll_Call) Return(_a0 []*models.Group, _a1 error) *GroupTable_GetAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDN provides a mock function with given fields: dn
func (_m *GroupTable) GetDN(dn string) (models.Group, error) {
	ret := _m.Called(dn)

	var r0 models.Group
	if rf, ok := ret.Get(0).(func(string) models.Group); ok {
		r0 = rf(dn)
	} else {
		r0 = ret.Get(0).(models.Group)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupTable_GetDN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDN'
type GroupTable_GetDN_Call struct {
	*mock.Call
}

// GetDN is a helper method to define mock.On call
//   - dn string
func (_e *GroupTable_Expecter) GetDN(dn interface{}) *GroupTable_GetDN_Call {
	return &GroupTable_GetDN_Call{Call: _e.mock.On("GetDN", dn)}
}

func (_c *GroupTable_GetDN_Call) Run(run func(dn string)) *GroupTable_GetDN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GroupTable_GetDN_Call) Return(_a0 models.Group, _a1 error) *GroupTable_GetDN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetID provides a mock function with given fields: id
func (_m *GroupTable) GetID(id int64) (models.Group, error) {
	ret := _m.Called(id)

	var r0 models.Group
	if rf, ok := ret.Get(0).(func(int64) models.Group); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Group)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GroupTable_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type GroupTable_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
//   - id int64
func (_e *GroupTable_Expecter) GetID(id interface{}) *GroupTable_GetID_Call {
	return &GroupTable_GetID_Call{Call: _e.mock.On("GetID", id)}
}

func (_c *GroupTable_GetID_Call) Run(run func(id int64)) *GroupTable_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *GroupTable_GetID_Call) Return(_a0 models.Group, _a1 error) *GroupTable_GetID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: group
func (_m *GroupTable) Insert(group *models.Group) error {
	ret := _m.Called(group)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Group) error); ok {
		r0 = rf(group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type GroupTable_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - group *models.Group
func (_e *GroupTable_Expecter) Insert(group interface{}) *GroupTable_Insert_Call {
	return &GroupTable_Insert_Call{Call: _e.mock.On("Insert", group)}
}

func (_c *GroupTable_Insert_Call) Run(run func(group *models.Group)) *GroupTable_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Group))
	})
	return _c
}

func (_c *GroupTable_Insert_Call) Return(_a0 error) *GroupTable_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// Update provides a mock function with given fields: group
func (_m *GroupTable) Update(group models.Group) error {
	ret := _m.Called(group)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Group) error); ok {
		r0 = rf(group)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type GroupTable_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - group models.Group
func (_e *GroupTable_Expecter) Update(group interface{}) *GroupTable_Update_Call {
	return &GroupTable_Update_Call{Call: _e.mock.On("Update", group)}
}

func (_c *GroupTable_Update_Call) Run(run func(group models.Group)) *GroupTable_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Group))
	})
	return _c
}

func (_c *GroupTable_Update_Call) Return(_a0 error) *GroupTable_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upgrade provides a mock function with given fields: old_version, new_verison
func (_m *GroupTable) Upgrade(old_version int, new_verison int) error {
	ret := _m.Called(old_version, new_verison)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(old_version, new_verison)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GroupTable_Upgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upgrade'
type GroupTable_Upgrade_Call struct {
	*mock.Call
}

// Upgrade is a helper method to define mock.On call
//   - old_version int
//   - new_verison int
func (_e *GroupTable_Expecter) Upgrade(old_version interface{}, new_verison interface{}) *GroupTable_Upgrade_Call {
	return &GroupTable_Upgrade_Call{Call: _e.mock.On("Upgrade", old_version, new_verison)}
}

func (_c *GroupTable_Upgrade_Call) Run(run func(old_version int, new_verison int)) *GroupTable_Upgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *GroupTable_Upgrade_Call) Return(_a0 error) *GroupTable_Upgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewGroupTable interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroupTable creates a new instance of GroupTable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroupTable(t mockConstructorTestingTNewGroupTable) *GroupTable {
	mock := &GroupTable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
