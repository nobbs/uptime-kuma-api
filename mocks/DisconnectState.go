// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DisconnectState is an autogenerated mock type for the DisconnectState type
type DisconnectState struct {
	mock.Mock
}

type DisconnectState_Expecter struct {
	mock *mock.Mock
}

func (_m *DisconnectState) EXPECT() *DisconnectState_Expecter {
	return &DisconnectState_Expecter{mock: &_m.Mock}
}

// SetConnected provides a mock function with given fields: _a0
func (_m *DisconnectState) SetConnected(_a0 bool) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DisconnectState_SetConnected_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConnected'
type DisconnectState_SetConnected_Call struct {
	*mock.Call
}

// SetConnected is a helper method to define mock.On call
//   - _a0 bool
func (_e *DisconnectState_Expecter) SetConnected(_a0 interface{}) *DisconnectState_SetConnected_Call {
	return &DisconnectState_SetConnected_Call{Call: _e.mock.On("SetConnected", _a0)}
}

func (_c *DisconnectState_SetConnected_Call) Run(run func(_a0 bool)) *DisconnectState_SetConnected_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return _c
}

func (_c *DisconnectState_SetConnected_Call) Return(_a0 error) *DisconnectState_SetConnected_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *DisconnectState_SetConnected_Call) RunAndReturn(run func(bool) error) *DisconnectState_SetConnected_Call {
	_c.Call.Return(run)
	return _c
}

// NewDisconnectState creates a new instance of DisconnectState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDisconnectState(t interface {
	mock.TestingT
	Cleanup(func())
}) *DisconnectState {
	mock := &DisconnectState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
