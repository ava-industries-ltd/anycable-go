// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// HandlePubSub provides a mock function with given fields: json
func (_m *Handler) HandlePubSub(json []byte) {
	_m.Called(json)
}
