// Code generated by mockery v1.0.0. DO NOT EDIT.

package specstack

import mock "github.com/stretchr/testify/mock"

// MockConfigAsserter is an autogenerated mock type for the ConfigAsserter type
type MockConfigAsserter struct {
	mock.Mock
}

// AssertConfig provides a mock function with given fields:
func (_m *MockConfigAsserter) AssertConfig() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
