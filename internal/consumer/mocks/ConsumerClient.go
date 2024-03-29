// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	sarama "github.com/Shopify/sarama"
	mock "github.com/stretchr/testify/mock"
)

// ConsumerClient is an autogenerated mock type for the ConsumerClient type
type ConsumerClient struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ConsumerClient) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ConsumePartition provides a mock function with given fields: topic, partition, offset
func (_m *ConsumerClient) ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error) {
	ret := _m.Called(topic, partition, offset)

	var r0 sarama.PartitionConsumer
	var r1 error
	if rf, ok := ret.Get(0).(func(string, int32, int64) (sarama.PartitionConsumer, error)); ok {
		return rf(topic, partition, offset)
	}
	if rf, ok := ret.Get(0).(func(string, int32, int64) sarama.PartitionConsumer); ok {
		r0 = rf(topic, partition, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sarama.PartitionConsumer)
		}
	}

	if rf, ok := ret.Get(1).(func(string, int32, int64) error); ok {
		r1 = rf(topic, partition, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Partitions provides a mock function with given fields: topic
func (_m *ConsumerClient) Partitions(topic string) ([]int32, error) {
	ret := _m.Called(topic)

	var r0 []int32
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]int32, error)); ok {
		return rf(topic)
	}
	if rf, ok := ret.Get(0).(func(string) []int32); ok {
		r0 = rf(topic)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int32)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(topic)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewConsumerClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsumerClient creates a new instance of ConsumerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsumerClient(t mockConstructorTestingTNewConsumerClient) *ConsumerClient {
	mock := &ConsumerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
