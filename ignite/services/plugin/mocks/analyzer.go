// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	v1 "github.com/ignite/cli/ignite/services/plugin/grpc/v1"
)

// PluginAnalyzer is an autogenerated mock type for the Analyzer type
type PluginAnalyzer struct {
	mock.Mock
}

type PluginAnalyzer_Expecter struct {
	mock *mock.Mock
}

func (_m *PluginAnalyzer) EXPECT() *PluginAnalyzer_Expecter {
	return &PluginAnalyzer_Expecter{mock: &_m.Mock}
}

// Dependencies provides a mock function with given fields: _a0
func (_m *PluginAnalyzer) Dependencies(_a0 context.Context) ([]*v1.Dependency, error) {
	ret := _m.Called(_a0)

	var r0 []*v1.Dependency
	if rf, ok := ret.Get(0).(func(context.Context) []*v1.Dependency); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1.Dependency)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PluginAnalyzer_Dependencies_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Dependencies'
type PluginAnalyzer_Dependencies_Call struct {
	*mock.Call
}

// Dependencies is a helper method to define mock.On call
//   - _a0 context.Context
func (_e *PluginAnalyzer_Expecter) Dependencies(_a0 interface{}) *PluginAnalyzer_Dependencies_Call {
	return &PluginAnalyzer_Dependencies_Call{Call: _e.mock.On("Dependencies", _a0)}
}

func (_c *PluginAnalyzer_Dependencies_Call) Run(run func(_a0 context.Context)) *PluginAnalyzer_Dependencies_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *PluginAnalyzer_Dependencies_Call) Return(_a0 []*v1.Dependency, _a1 error) *PluginAnalyzer_Dependencies_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewPluginAnalyzer interface {
	mock.TestingT
	Cleanup(func())
}

// NewPluginAnalyzer creates a new instance of PluginAnalyzer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPluginAnalyzer(t mockConstructorTestingTNewPluginAnalyzer) *PluginAnalyzer {
	mock := &PluginAnalyzer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}