// Code generated by mockery v2.18.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

// ApplicationsTable is an autogenerated mock type for the ApplicationsTable type
type ApplicationsTable struct {
	mock.Mock
}

type ApplicationsTable_Expecter struct {
	mock *mock.Mock
}

func (_m *ApplicationsTable) EXPECT() *ApplicationsTable_Expecter {
	return &ApplicationsTable_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields:
func (_m *ApplicationsTable) Create() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type ApplicationsTable_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
func (_e *ApplicationsTable_Expecter) Create() *ApplicationsTable_Create_Call {
	return &ApplicationsTable_Create_Call{Call: _e.mock.On("Create")}
}

func (_c *ApplicationsTable_Create_Call) Run(run func()) *ApplicationsTable_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ApplicationsTable_Create_Call) Return(_a0 error) *ApplicationsTable_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

// Delete provides a mock function with given fields: id
func (_m *ApplicationsTable) Delete(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type ApplicationsTable_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - id int64
func (_e *ApplicationsTable_Expecter) Delete(id interface{}) *ApplicationsTable_Delete_Call {
	return &ApplicationsTable_Delete_Call{Call: _e.mock.On("Delete", id)}
}

func (_c *ApplicationsTable_Delete_Call) Run(run func(id int64)) *ApplicationsTable_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ApplicationsTable_Delete_Call) Return(_a0 error) *ApplicationsTable_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

// Drop provides a mock function with given fields:
func (_m *ApplicationsTable) Drop() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Drop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Drop'
type ApplicationsTable_Drop_Call struct {
	*mock.Call
}

// Drop is a helper method to define mock.On call
func (_e *ApplicationsTable_Expecter) Drop() *ApplicationsTable_Drop_Call {
	return &ApplicationsTable_Drop_Call{Call: _e.mock.On("Drop")}
}

func (_c *ApplicationsTable_Drop_Call) Run(run func()) *ApplicationsTable_Drop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ApplicationsTable_Drop_Call) Return(_a0 error) *ApplicationsTable_Drop_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetID provides a mock function with given fields: id
func (_m *ApplicationsTable) GetID(id int64) (models.Application, error) {
	ret := _m.Called(id)

	var r0 models.Application
	if rf, ok := ret.Get(0).(func(int64) models.Application); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Application)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplicationsTable_GetID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetID'
type ApplicationsTable_GetID_Call struct {
	*mock.Call
}

// GetID is a helper method to define mock.On call
//   - id int64
func (_e *ApplicationsTable_Expecter) GetID(id interface{}) *ApplicationsTable_GetID_Call {
	return &ApplicationsTable_GetID_Call{Call: _e.mock.On("GetID", id)}
}

func (_c *ApplicationsTable_GetID_Call) Run(run func(id int64)) *ApplicationsTable_GetID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *ApplicationsTable_GetID_Call) Return(_a0 models.Application, _a1 error) *ApplicationsTable_GetID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTemplateID provides a mock function with given fields: appid
func (_m *ApplicationsTable) GetTemplateID(appid string) ([]*models.Application, error) {
	ret := _m.Called(appid)

	var r0 []*models.Application
	if rf, ok := ret.Get(0).(func(string) []*models.Application); ok {
		r0 = rf(appid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Application)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(appid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApplicationsTable_GetTemplateID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplateID'
type ApplicationsTable_GetTemplateID_Call struct {
	*mock.Call
}

// GetTemplateID is a helper method to define mock.On call
//   - appid string
func (_e *ApplicationsTable_Expecter) GetTemplateID(appid interface{}) *ApplicationsTable_GetTemplateID_Call {
	return &ApplicationsTable_GetTemplateID_Call{Call: _e.mock.On("GetTemplateID", appid)}
}

func (_c *ApplicationsTable_GetTemplateID_Call) Run(run func(appid string)) *ApplicationsTable_GetTemplateID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ApplicationsTable_GetTemplateID_Call) Return(_a0 []*models.Application, _a1 error) *ApplicationsTable_GetTemplateID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Insert provides a mock function with given fields: app
func (_m *ApplicationsTable) Insert(app *models.Application) error {
	ret := _m.Called(app)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.Application) error); ok {
		r0 = rf(app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Insert_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Insert'
type ApplicationsTable_Insert_Call struct {
	*mock.Call
}

// Insert is a helper method to define mock.On call
//   - app *models.Application
func (_e *ApplicationsTable_Expecter) Insert(app interface{}) *ApplicationsTable_Insert_Call {
	return &ApplicationsTable_Insert_Call{Call: _e.mock.On("Insert", app)}
}

func (_c *ApplicationsTable_Insert_Call) Run(run func(app *models.Application)) *ApplicationsTable_Insert_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*models.Application))
	})
	return _c
}

func (_c *ApplicationsTable_Insert_Call) Return(_a0 error) *ApplicationsTable_Insert_Call {
	_c.Call.Return(_a0)
	return _c
}

// Update provides a mock function with given fields: app
func (_m *ApplicationsTable) Update(app models.Application) error {
	ret := _m.Called(app)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Application) error); ok {
		r0 = rf(app)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type ApplicationsTable_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - app models.Application
func (_e *ApplicationsTable_Expecter) Update(app interface{}) *ApplicationsTable_Update_Call {
	return &ApplicationsTable_Update_Call{Call: _e.mock.On("Update", app)}
}

func (_c *ApplicationsTable_Update_Call) Run(run func(app models.Application)) *ApplicationsTable_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Application))
	})
	return _c
}

func (_c *ApplicationsTable_Update_Call) Return(_a0 error) *ApplicationsTable_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

// Upgrade provides a mock function with given fields: old_version, new_verison
func (_m *ApplicationsTable) Upgrade(old_version int, new_verison int) error {
	ret := _m.Called(old_version, new_verison)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, int) error); ok {
		r0 = rf(old_version, new_verison)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApplicationsTable_Upgrade_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upgrade'
type ApplicationsTable_Upgrade_Call struct {
	*mock.Call
}

// Upgrade is a helper method to define mock.On call
//   - old_version int
//   - new_verison int
func (_e *ApplicationsTable_Expecter) Upgrade(old_version interface{}, new_verison interface{}) *ApplicationsTable_Upgrade_Call {
	return &ApplicationsTable_Upgrade_Call{Call: _e.mock.On("Upgrade", old_version, new_verison)}
}

func (_c *ApplicationsTable_Upgrade_Call) Run(run func(old_version int, new_verison int)) *ApplicationsTable_Upgrade_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *ApplicationsTable_Upgrade_Call) Return(_a0 error) *ApplicationsTable_Upgrade_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewApplicationsTable interface {
	mock.TestingT
	Cleanup(func())
}

// NewApplicationsTable creates a new instance of ApplicationsTable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApplicationsTable(t mockConstructorTestingTNewApplicationsTable) *ApplicationsTable {
	mock := &ApplicationsTable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
