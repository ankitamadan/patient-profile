// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	sarama "github.com/Shopify/sarama"
	mock "github.com/stretchr/testify/mock"
)

// PartitionConsumerClient is an autogenerated mock type for the PartitionConsumerClient type
type PartitionConsumerClient struct {
	mock.Mock
}

// AsyncClose provides a mock function with given fields:
func (_m *PartitionConsumerClient) AsyncClose() {
	_m.Called()
}

// Close provides a mock function with given fields:
func (_m *PartitionConsumerClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Errors provides a mock function with given fields:
func (_m *PartitionConsumerClient) Errors() <-chan *sarama.ConsumerError {
	ret := _m.Called()

	var r0 <-chan *sarama.ConsumerError
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ConsumerError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ConsumerError)
		}
	}

	return r0
}

// HighWaterMarkOffset provides a mock function with given fields:
func (_m *PartitionConsumerClient) HighWaterMarkOffset() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// Messages provides a mock function with given fields:
func (_m *PartitionConsumerClient) Messages() <-chan *sarama.ConsumerMessage {
	ret := _m.Called()

	var r0 <-chan *sarama.ConsumerMessage
	if rf, ok := ret.Get(0).(func() <-chan *sarama.ConsumerMessage); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *sarama.ConsumerMessage)
		}
	}

	return r0
}

type mockConstructorTestingTNewPartitionConsumerClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewPartitionConsumerClient creates a new instance of PartitionConsumerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPartitionConsumerClient(t mockConstructorTestingTNewPartitionConsumerClient) *PartitionConsumerClient {
	mock := &PartitionConsumerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
