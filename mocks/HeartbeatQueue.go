// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	state "github.com/nobbs/uptime-kuma-api/pkg/state"
	mock "github.com/stretchr/testify/mock"
)

// HeartbeatQueue is an autogenerated mock type for the HeartbeatQueue type
type HeartbeatQueue struct {
	mock.Mock
}

type HeartbeatQueue_Expecter struct {
	mock *mock.Mock
}

func (_m *HeartbeatQueue) EXPECT() *HeartbeatQueue_Expecter {
	return &HeartbeatQueue_Expecter{mock: &_m.Mock}
}

// Push provides a mock function with given fields: beat
func (_m *HeartbeatQueue) Push(beat *state.Heartbeat) {
	_m.Called(beat)
}

// HeartbeatQueue_Push_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Push'
type HeartbeatQueue_Push_Call struct {
	*mock.Call
}

// Push is a helper method to define mock.On call
//   - beat *state.Heartbeat
func (_e *HeartbeatQueue_Expecter) Push(beat interface{}) *HeartbeatQueue_Push_Call {
	return &HeartbeatQueue_Push_Call{Call: _e.mock.On("Push", beat)}
}

func (_c *HeartbeatQueue_Push_Call) Run(run func(beat *state.Heartbeat)) *HeartbeatQueue_Push_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*state.Heartbeat))
	})
	return _c
}

func (_c *HeartbeatQueue_Push_Call) Return() *HeartbeatQueue_Push_Call {
	_c.Call.Return()
	return _c
}

func (_c *HeartbeatQueue_Push_Call) RunAndReturn(run func(*state.Heartbeat)) *HeartbeatQueue_Push_Call {
	_c.Call.Return(run)
	return _c
}

// Slice provides a mock function with given fields:
func (_m *HeartbeatQueue) Slice() []state.Heartbeat {
	ret := _m.Called()

	var r0 []state.Heartbeat
	if rf, ok := ret.Get(0).(func() []state.Heartbeat); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]state.Heartbeat)
		}
	}

	return r0
}

// HeartbeatQueue_Slice_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Slice'
type HeartbeatQueue_Slice_Call struct {
	*mock.Call
}

// Slice is a helper method to define mock.On call
func (_e *HeartbeatQueue_Expecter) Slice() *HeartbeatQueue_Slice_Call {
	return &HeartbeatQueue_Slice_Call{Call: _e.mock.On("Slice")}
}

func (_c *HeartbeatQueue_Slice_Call) Run(run func()) *HeartbeatQueue_Slice_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *HeartbeatQueue_Slice_Call) Return(_a0 []state.Heartbeat) *HeartbeatQueue_Slice_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *HeartbeatQueue_Slice_Call) RunAndReturn(run func() []state.Heartbeat) *HeartbeatQueue_Slice_Call {
	_c.Call.Return(run)
	return _c
}

// Trim provides a mock function with given fields: capacity
func (_m *HeartbeatQueue) Trim(capacity int) {
	_m.Called(capacity)
}

// HeartbeatQueue_Trim_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Trim'
type HeartbeatQueue_Trim_Call struct {
	*mock.Call
}

// Trim is a helper method to define mock.On call
//   - capacity int
func (_e *HeartbeatQueue_Expecter) Trim(capacity interface{}) *HeartbeatQueue_Trim_Call {
	return &HeartbeatQueue_Trim_Call{Call: _e.mock.On("Trim", capacity)}
}

func (_c *HeartbeatQueue_Trim_Call) Run(run func(capacity int)) *HeartbeatQueue_Trim_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *HeartbeatQueue_Trim_Call) Return() *HeartbeatQueue_Trim_Call {
	_c.Call.Return()
	return _c
}

func (_c *HeartbeatQueue_Trim_Call) RunAndReturn(run func(int)) *HeartbeatQueue_Trim_Call {
	_c.Call.Return(run)
	return _c
}

// NewHeartbeatQueue creates a new instance of HeartbeatQueue. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHeartbeatQueue(t interface {
	mock.TestingT
	Cleanup(func())
}) *HeartbeatQueue {
	mock := &HeartbeatQueue{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
