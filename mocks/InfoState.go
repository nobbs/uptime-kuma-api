// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	state "github.com/nobbs/uptime-kuma-api/pkg/state"
	mock "github.com/stretchr/testify/mock"
)

// InfoState is an autogenerated mock type for the InfoState type
type InfoState struct {
	mock.Mock
}

type InfoState_Expecter struct {
	mock *mock.Mock
}

func (_m *InfoState) EXPECT() *InfoState_Expecter {
	return &InfoState_Expecter{mock: &_m.Mock}
}

// Info provides a mock function with given fields:
func (_m *InfoState) Info() (*state.Info, error) {
	ret := _m.Called()

	var r0 *state.Info
	var r1 error
	if rf, ok := ret.Get(0).(func() (*state.Info, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() *state.Info); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*state.Info)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InfoState_Info_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Info'
type InfoState_Info_Call struct {
	*mock.Call
}

// Info is a helper method to define mock.On call
func (_e *InfoState_Expecter) Info() *InfoState_Info_Call {
	return &InfoState_Info_Call{Call: _e.mock.On("Info")}
}

func (_c *InfoState_Info_Call) Run(run func()) *InfoState_Info_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *InfoState_Info_Call) Return(_a0 *state.Info, _a1 error) *InfoState_Info_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *InfoState_Info_Call) RunAndReturn(run func() (*state.Info, error)) *InfoState_Info_Call {
	_c.Call.Return(run)
	return _c
}

// SetInfo provides a mock function with given fields: _a0
func (_m *InfoState) SetInfo(_a0 *state.Info) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*state.Info) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InfoState_SetInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetInfo'
type InfoState_SetInfo_Call struct {
	*mock.Call
}

// SetInfo is a helper method to define mock.On call
//   - _a0 *state.Info
func (_e *InfoState_Expecter) SetInfo(_a0 interface{}) *InfoState_SetInfo_Call {
	return &InfoState_SetInfo_Call{Call: _e.mock.On("SetInfo", _a0)}
}

func (_c *InfoState_SetInfo_Call) Run(run func(_a0 *state.Info)) *InfoState_SetInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*state.Info))
	})
	return _c
}

func (_c *InfoState_SetInfo_Call) Return(_a0 error) *InfoState_SetInfo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *InfoState_SetInfo_Call) RunAndReturn(run func(*state.Info) error) *InfoState_SetInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewInfoState creates a new instance of InfoState. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInfoState(t interface {
	mock.TestingT
	Cleanup(func())
}) *InfoState {
	mock := &InfoState{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
