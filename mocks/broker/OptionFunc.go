// Code generated by mockery v2.8.0. DO NOT EDIT.

package mocks

import (
	broker "github.com/golangid/candi/broker"
	mock "github.com/stretchr/testify/mock"
)

// OptionFunc is an autogenerated mock type for the OptionFunc type
type OptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: _a0
func (_m *OptionFunc) Execute(_a0 *broker.Broker) {
	_m.Called(_a0)
}
