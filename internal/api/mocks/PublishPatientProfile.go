// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PublishPatientProfile is an autogenerated mock type for the PublishPatientProfile type
type PublishPatientProfile struct {
	mock.Mock
}

// PushPatientProfileToQueue provides a mock function with given fields: message
func (_m *PublishPatientProfile) PushPatientProfileToQueue(message []byte) error {
	ret := _m.Called(message)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewPublishPatientProfile interface {
	mock.TestingT
	Cleanup(func())
}

// NewPublishPatientProfile creates a new instance of PublishPatientProfile. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPublishPatientProfile(t mockConstructorTestingTNewPublishPatientProfile) *PublishPatientProfile {
	mock := &PublishPatientProfile{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
