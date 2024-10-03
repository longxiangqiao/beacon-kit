// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Dispatcher is an autogenerated mock type for the Dispatcher type
type Dispatcher struct {
	mock.Mock
}

type Dispatcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Dispatcher) EXPECT() *Dispatcher_Expecter {
	return &Dispatcher_Expecter{mock: &_m.Mock}
}

// Start provides a mock function with given fields: ctx
func (_m *Dispatcher) Start(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dispatcher_Start_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Start'
type Dispatcher_Start_Call struct {
	*mock.Call
}

// Start is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Dispatcher_Expecter) Start(ctx interface{}) *Dispatcher_Start_Call {
	return &Dispatcher_Start_Call{Call: _e.mock.On("Start", ctx)}
}

func (_c *Dispatcher_Start_Call) Run(run func(ctx context.Context)) *Dispatcher_Start_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Dispatcher_Start_Call) Return(_a0 error) *Dispatcher_Start_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Dispatcher_Start_Call) RunAndReturn(run func(context.Context) error) *Dispatcher_Start_Call {
	_c.Call.Return(run)
	return _c
}

// NewDispatcher creates a new instance of Dispatcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDispatcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Dispatcher {
	mock := &Dispatcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
