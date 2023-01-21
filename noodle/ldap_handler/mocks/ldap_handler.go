// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	models "github.com/mwinters-stuff/noodle/server/models"
	mock "github.com/stretchr/testify/mock"
)

// LdapHandler is an autogenerated mock type for the LdapHandler type
type LdapHandler struct {
	mock.Mock
}

type LdapHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *LdapHandler) EXPECT() *LdapHandler_Expecter {
	return &LdapHandler_Expecter{mock: &_m.Mock}
}

// AuthUser provides a mock function with given fields: userdn, password
func (_m *LdapHandler) AuthUser(userdn string, password string) (bool, error) {
	ret := _m.Called(userdn, password)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(userdn, password)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userdn, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LdapHandler_AuthUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AuthUser'
type LdapHandler_AuthUser_Call struct {
	*mock.Call
}

// AuthUser is a helper method to define mock.On call
//   - userdn string
//   - password string
func (_e *LdapHandler_Expecter) AuthUser(userdn interface{}, password interface{}) *LdapHandler_AuthUser_Call {
	return &LdapHandler_AuthUser_Call{Call: _e.mock.On("AuthUser", userdn, password)}
}

func (_c *LdapHandler_AuthUser_Call) Run(run func(userdn string, password string)) *LdapHandler_AuthUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *LdapHandler_AuthUser_Call) Return(_a0 bool, _a1 error) *LdapHandler_AuthUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// Connect provides a mock function with given fields:
func (_m *LdapHandler) Connect() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LdapHandler_Connect_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Connect'
type LdapHandler_Connect_Call struct {
	*mock.Call
}

// Connect is a helper method to define mock.On call
func (_e *LdapHandler_Expecter) Connect() *LdapHandler_Connect_Call {
	return &LdapHandler_Connect_Call{Call: _e.mock.On("Connect")}
}

func (_c *LdapHandler_Connect_Call) Run(run func()) *LdapHandler_Connect_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *LdapHandler_Connect_Call) Return(_a0 error) *LdapHandler_Connect_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetGroupUsers provides a mock function with given fields: _a0
func (_m *LdapHandler) GetGroupUsers(_a0 models.Group) ([]models.UserGroup, error) {
	ret := _m.Called(_a0)

	var r0 []models.UserGroup
	if rf, ok := ret.Get(0).(func(models.Group) []models.UserGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.UserGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.Group) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LdapHandler_GetGroupUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroupUsers'
type LdapHandler_GetGroupUsers_Call struct {
	*mock.Call
}

// GetGroupUsers is a helper method to define mock.On call
//   - _a0 models.Group
func (_e *LdapHandler_Expecter) GetGroupUsers(_a0 interface{}) *LdapHandler_GetGroupUsers_Call {
	return &LdapHandler_GetGroupUsers_Call{Call: _e.mock.On("GetGroupUsers", _a0)}
}

func (_c *LdapHandler_GetGroupUsers_Call) Run(run func(_a0 models.Group)) *LdapHandler_GetGroupUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.Group))
	})
	return _c
}

func (_c *LdapHandler_GetGroupUsers_Call) Return(_a0 []models.UserGroup, _a1 error) *LdapHandler_GetGroupUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetGroups provides a mock function with given fields:
func (_m *LdapHandler) GetGroups() ([]models.Group, error) {
	ret := _m.Called()

	var r0 []models.Group
	if rf, ok := ret.Get(0).(func() []models.Group); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Group)
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

// LdapHandler_GetGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetGroups'
type LdapHandler_GetGroups_Call struct {
	*mock.Call
}

// GetGroups is a helper method to define mock.On call
func (_e *LdapHandler_Expecter) GetGroups() *LdapHandler_GetGroups_Call {
	return &LdapHandler_GetGroups_Call{Call: _e.mock.On("GetGroups")}
}

func (_c *LdapHandler_GetGroups_Call) Run(run func()) *LdapHandler_GetGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *LdapHandler_GetGroups_Call) Return(_a0 []models.Group, _a1 error) *LdapHandler_GetGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUser provides a mock function with given fields: username
func (_m *LdapHandler) GetUser(username string) (models.User, error) {
	ret := _m.Called(username)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LdapHandler_GetUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUser'
type LdapHandler_GetUser_Call struct {
	*mock.Call
}

// GetUser is a helper method to define mock.On call
//   - username string
func (_e *LdapHandler_Expecter) GetUser(username interface{}) *LdapHandler_GetUser_Call {
	return &LdapHandler_GetUser_Call{Call: _e.mock.On("GetUser", username)}
}

func (_c *LdapHandler_GetUser_Call) Run(run func(username string)) *LdapHandler_GetUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LdapHandler_GetUser_Call) Return(_a0 models.User, _a1 error) *LdapHandler_GetUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUserByDN provides a mock function with given fields: dn
func (_m *LdapHandler) GetUserByDN(dn string) (models.User, error) {
	ret := _m.Called(dn)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(dn)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(dn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LdapHandler_GetUserByDN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByDN'
type LdapHandler_GetUserByDN_Call struct {
	*mock.Call
}

// GetUserByDN is a helper method to define mock.On call
//   - dn string
func (_e *LdapHandler_Expecter) GetUserByDN(dn interface{}) *LdapHandler_GetUserByDN_Call {
	return &LdapHandler_GetUserByDN_Call{Call: _e.mock.On("GetUserByDN", dn)}
}

func (_c *LdapHandler_GetUserByDN_Call) Run(run func(dn string)) *LdapHandler_GetUserByDN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *LdapHandler_GetUserByDN_Call) Return(_a0 models.User, _a1 error) *LdapHandler_GetUserByDN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUserGroups provides a mock function with given fields: _a0
func (_m *LdapHandler) GetUserGroups(_a0 models.User) ([]models.UserGroup, error) {
	ret := _m.Called(_a0)

	var r0 []models.UserGroup
	if rf, ok := ret.Get(0).(func(models.User) []models.UserGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.UserGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(models.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LdapHandler_GetUserGroups_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserGroups'
type LdapHandler_GetUserGroups_Call struct {
	*mock.Call
}

// GetUserGroups is a helper method to define mock.On call
//   - _a0 models.User
func (_e *LdapHandler_Expecter) GetUserGroups(_a0 interface{}) *LdapHandler_GetUserGroups_Call {
	return &LdapHandler_GetUserGroups_Call{Call: _e.mock.On("GetUserGroups", _a0)}
}

func (_c *LdapHandler_GetUserGroups_Call) Run(run func(_a0 models.User)) *LdapHandler_GetUserGroups_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(models.User))
	})
	return _c
}

func (_c *LdapHandler_GetUserGroups_Call) Return(_a0 []models.UserGroup, _a1 error) *LdapHandler_GetUserGroups_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetUsers provides a mock function with given fields:
func (_m *LdapHandler) GetUsers() ([]models.User, error) {
	ret := _m.Called()

	var r0 []models.User
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
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

// LdapHandler_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type LdapHandler_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
func (_e *LdapHandler_Expecter) GetUsers() *LdapHandler_GetUsers_Call {
	return &LdapHandler_GetUsers_Call{Call: _e.mock.On("GetUsers")}
}

func (_c *LdapHandler_GetUsers_Call) Run(run func()) *LdapHandler_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *LdapHandler_GetUsers_Call) Return(_a0 []models.User, _a1 error) *LdapHandler_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewLdapHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewLdapHandler creates a new instance of LdapHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLdapHandler(t mockConstructorTestingTNewLdapHandler) *LdapHandler {
	mock := &LdapHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}