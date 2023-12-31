// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	state "github.com/nobbs/uptime-kuma-api/pkg/state"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// StatefulEmiter is an autogenerated mock type for the StatefulEmiter type
type StatefulEmiter struct {
	mock.Mock
}

type StatefulEmiter_Expecter struct {
	mock *mock.Mock
}

func (_m *StatefulEmiter) EXPECT() *StatefulEmiter_Expecter {
	return &StatefulEmiter_Expecter{mock: &_m.Mock}
}

// Await provides a mock function with given fields: _a0, _a1
func (_m *StatefulEmiter) Await(_a0 string, _a1 time.Duration) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, time.Duration) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StatefulEmiter_Await_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Await'
type StatefulEmiter_Await_Call struct {
	*mock.Call
}

// Await is a helper method to define mock.On call
//   - _a0 string
//   - _a1 time.Duration
func (_e *StatefulEmiter_Expecter) Await(_a0 interface{}, _a1 interface{}) *StatefulEmiter_Await_Call {
	return &StatefulEmiter_Await_Call{Call: _e.mock.On("Await", _a0, _a1)}
}

func (_c *StatefulEmiter_Await_Call) Run(run func(_a0 string, _a1 time.Duration)) *StatefulEmiter_Await_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(time.Duration))
	})
	return _c
}

func (_c *StatefulEmiter_Await_Call) Return(_a0 error) *StatefulEmiter_Await_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StatefulEmiter_Await_Call) RunAndReturn(run func(string, time.Duration) error) *StatefulEmiter_Await_Call {
	_c.Call.Return(run)
	return _c
}

// Emit provides a mock function with given fields: _a0, _a1, _a2
func (_m *StatefulEmiter) Emit(_a0 string, _a1 time.Duration, _a2 ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _a2...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, time.Duration, ...interface{}) (interface{}, error)); ok {
		return rf(_a0, _a1, _a2...)
	}
	if rf, ok := ret.Get(0).(func(string, time.Duration, ...interface{}) interface{}); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, time.Duration, ...interface{}) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StatefulEmiter_Emit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Emit'
type StatefulEmiter_Emit_Call struct {
	*mock.Call
}

// Emit is a helper method to define mock.On call
//   - _a0 string
//   - _a1 time.Duration
//   - _a2 ...interface{}
func (_e *StatefulEmiter_Expecter) Emit(_a0 interface{}, _a1 interface{}, _a2 ...interface{}) *StatefulEmiter_Emit_Call {
	return &StatefulEmiter_Emit_Call{Call: _e.mock.On("Emit",
		append([]interface{}{_a0, _a1}, _a2...)...)}
}

func (_c *StatefulEmiter_Emit_Call) Run(run func(_a0 string, _a1 time.Duration, _a2 ...interface{})) *StatefulEmiter_Emit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]interface{}, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(interface{})
			}
		}
		run(args[0].(string), args[1].(time.Duration), variadicArgs...)
	})
	return _c
}

func (_c *StatefulEmiter_Emit_Call) Return(_a0 interface{}, _a1 error) *StatefulEmiter_Emit_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StatefulEmiter_Emit_Call) RunAndReturn(run func(string, time.Duration, ...interface{}) (interface{}, error)) *StatefulEmiter_Emit_Call {
	_c.Call.Return(run)
	return _c
}

// State provides a mock function with given fields:
func (_m *StatefulEmiter) State() *state.State {
	ret := _m.Called()

	var r0 *state.State
	if rf, ok := ret.Get(0).(func() *state.State); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.State)
		}
	}

	return r0
}

// StatefulEmiter_State_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'State'
type StatefulEmiter_State_Call struct {
	*mock.Call
}

// State is a helper method to define mock.On call
func (_e *StatefulEmiter_Expecter) State() *StatefulEmiter_State_Call {
	return &StatefulEmiter_State_Call{Call: _e.mock.On("State")}
}

func (_c *StatefulEmiter_State_Call) Run(run func()) *StatefulEmiter_State_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *StatefulEmiter_State_Call) Return(_a0 *state.State) *StatefulEmiter_State_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *StatefulEmiter_State_Call) RunAndReturn(run func() *state.State) *StatefulEmiter_State_Call {
	_c.Call.Return(run)
	return _c
}

// NewStatefulEmiter creates a new instance of StatefulEmiter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStatefulEmiter(t interface {
	mock.TestingT
	Cleanup(func())
}) *StatefulEmiter {
	mock := &StatefulEmiter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
