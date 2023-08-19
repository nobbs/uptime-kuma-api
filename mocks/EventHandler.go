// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	handler "github.com/nobbs/uptime-kuma-api/pkg/handler"
	mock "github.com/stretchr/testify/mock"
)

// EventHandler is an autogenerated mock type for the EventHandler type
type EventHandler struct {
	mock.Mock
}

type EventHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *EventHandler) EXPECT() *EventHandler_Expecter {
	return &EventHandler_Expecter{mock: &_m.Mock}
}

// Event provides a mock function with given fields:
func (_m *EventHandler) Event() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// EventHandler_Event_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Event'
type EventHandler_Event_Call struct {
	*mock.Call
}

// Event is a helper method to define mock.On call
func (_e *EventHandler_Expecter) Event() *EventHandler_Event_Call {
	return &EventHandler_Event_Call{Call: _e.mock.On("Event")}
}

func (_c *EventHandler_Event_Call) Run(run func()) *EventHandler_Event_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EventHandler_Event_Call) Return(_a0 string) *EventHandler_Event_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventHandler_Event_Call) RunAndReturn(run func() string) *EventHandler_Event_Call {
	_c.Call.Return(run)
	return _c
}

// Occured provides a mock function with given fields:
func (_m *EventHandler) Occured() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EventHandler_Occured_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Occured'
type EventHandler_Occured_Call struct {
	*mock.Call
}

// Occured is a helper method to define mock.On call
func (_e *EventHandler_Expecter) Occured() *EventHandler_Occured_Call {
	return &EventHandler_Occured_Call{Call: _e.mock.On("Occured")}
}

func (_c *EventHandler_Occured_Call) Run(run func()) *EventHandler_Occured_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *EventHandler_Occured_Call) Return(_a0 bool) *EventHandler_Occured_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventHandler_Occured_Call) RunAndReturn(run func() bool) *EventHandler_Occured_Call {
	_c.Call.Return(run)
	return _c
}

// Register provides a mock function with given fields: _a0
func (_m *EventHandler) Register(_a0 handler.HandlerRegistrator) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(handler.HandlerRegistrator) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EventHandler_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type EventHandler_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 handler.HandlerRegistrator
func (_e *EventHandler_Expecter) Register(_a0 interface{}) *EventHandler_Register_Call {
	return &EventHandler_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *EventHandler_Register_Call) Run(run func(_a0 handler.HandlerRegistrator)) *EventHandler_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(handler.HandlerRegistrator))
	})
	return _c
}

func (_c *EventHandler_Register_Call) Return(_a0 error) *EventHandler_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *EventHandler_Register_Call) RunAndReturn(run func(handler.HandlerRegistrator) error) *EventHandler_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewEventHandler creates a new instance of EventHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewEventHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *EventHandler {
	mock := &EventHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
