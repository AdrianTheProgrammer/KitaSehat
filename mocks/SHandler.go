// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// SHandler is an autogenerated mock type for the SHandler type
type SHandler struct {
	mock.Mock
}

// AddSurvey provides a mock function with given fields: _a0
func (_m *SHandler) AddSurvey(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for AddSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSurvey provides a mock function with given fields: _a0
func (_m *SHandler) DeleteSurvey(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllSurveys provides a mock function with given fields: _a0
func (_m *SHandler) GetAllSurveys(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetAllSurveys")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetSurvey provides a mock function with given fields: _a0
func (_m *SHandler) GetSurvey(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSurvey provides a mock function with given fields: _a0
func (_m *SHandler) UpdateSurvey(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateSurvey")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewSHandler creates a new instance of SHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *SHandler {
	mock := &SHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
