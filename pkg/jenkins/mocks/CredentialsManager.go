// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CredentialsManager is an autogenerated mock type for the CredentialsManager type
type CredentialsManager struct {
	mock.Mock
}

// Update provides a mock function with given fields: _a0, _a1, _a2
func (_m *CredentialsManager) Update(_a0 string, _a1 string, _a2 interface{}) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, interface{}) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}