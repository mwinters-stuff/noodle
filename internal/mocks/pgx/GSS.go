// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// GSS is an autogenerated mock type for the GSS type
type GSS struct {
	mock.Mock
}

type GSS_Expecter struct {
	mock *mock.Mock
}

func (_m *GSS) EXPECT() *GSS_Expecter {
	return &GSS_Expecter{mock: &_m.Mock}
}

// Continue provides a mock function with given fields: inToken
func (_m *GSS) Continue(inToken []byte) (bool, []byte, error) {
	ret := _m.Called(inToken)

	var r0 bool
	if rf, ok := ret.Get(0).(func([]byte) bool); ok {
		r0 = rf(inToken)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func([]byte) []byte); ok {
		r1 = rf(inToken)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]byte) error); ok {
		r2 = rf(inToken)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GSS_Continue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Continue'
type GSS_Continue_Call struct {
	*mock.Call
}

// Continue is a helper method to define mock.On call
//   - inToken []byte
func (_e *GSS_Expecter) Continue(inToken interface{}) *GSS_Continue_Call {
	return &GSS_Continue_Call{Call: _e.mock.On("Continue", inToken)}
}

func (_c *GSS_Continue_Call) Run(run func(inToken []byte)) *GSS_Continue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte))
	})
	return _c
}

func (_c *GSS_Continue_Call) Return(done bool, outToken []byte, err error) *GSS_Continue_Call {
	_c.Call.Return(done, outToken, err)
	return _c
}

// GetInitToken provides a mock function with given fields: host, service
func (_m *GSS) GetInitToken(host string, service string) ([]byte, error) {
	ret := _m.Called(host, service)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, string) []byte); ok {
		r0 = rf(host, service)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(host, service)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GSS_GetInitToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInitToken'
type GSS_GetInitToken_Call struct {
	*mock.Call
}

// GetInitToken is a helper method to define mock.On call
//   - host string
//   - service string
func (_e *GSS_Expecter) GetInitToken(host interface{}, service interface{}) *GSS_GetInitToken_Call {
	return &GSS_GetInitToken_Call{Call: _e.mock.On("GetInitToken", host, service)}
}

func (_c *GSS_GetInitToken_Call) Run(run func(host string, service string)) *GSS_GetInitToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *GSS_GetInitToken_Call) Return(_a0 []byte, _a1 error) *GSS_GetInitToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetInitTokenFromSPN provides a mock function with given fields: spn
func (_m *GSS) GetInitTokenFromSPN(spn string) ([]byte, error) {
	ret := _m.Called(spn)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(spn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(spn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GSS_GetInitTokenFromSPN_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetInitTokenFromSPN'
type GSS_GetInitTokenFromSPN_Call struct {
	*mock.Call
}

// GetInitTokenFromSPN is a helper method to define mock.On call
//   - spn string
func (_e *GSS_Expecter) GetInitTokenFromSPN(spn interface{}) *GSS_GetInitTokenFromSPN_Call {
	return &GSS_GetInitTokenFromSPN_Call{Call: _e.mock.On("GetInitTokenFromSPN", spn)}
}

func (_c *GSS_GetInitTokenFromSPN_Call) Run(run func(spn string)) *GSS_GetInitTokenFromSPN_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *GSS_GetInitTokenFromSPN_Call) Return(_a0 []byte, _a1 error) *GSS_GetInitTokenFromSPN_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewGSS interface {
	mock.TestingT
	Cleanup(func())
}

// NewGSS creates a new instance of GSS. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGSS(t mockConstructorTestingTNewGSS) *GSS {
	mock := &GSS{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
