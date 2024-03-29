// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"
	store "heidi_health/internal/store"

	mock "github.com/stretchr/testify/mock"
)

// StoreWriter is an autogenerated mock type for the StoreWriter type
type StoreWriter struct {
	mock.Mock
}

// InsertPatientProfile provides a mock function with given fields: _a0, _a1
func (_m *StoreWriter) InsertPatientProfile(_a0 context.Context, _a1 store.PatientProfile) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, store.PatientProfile) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewStoreWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewStoreWriter creates a new instance of StoreWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStoreWriter(t mockConstructorTestingTNewStoreWriter) *StoreWriter {
	mock := &StoreWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
