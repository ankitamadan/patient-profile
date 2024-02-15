package consumer_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Shopify/sarama"

	"heidi_health/internal/consumer"
	"heidi_health/internal/consumer/mocks"
	ddbMock "heidi_health/internal/store/mocks"
)

func TestConsumeMessage(t *testing.T) {

	mockWriter := ddbMock.NewStoreWriter(t)
	mockConsumer := mocks.NewConsumerClient(t)
	mockPartitionConsumerClient := mocks.NewPartitionConsumerClient(t)
	c := consumer.NewConsumer("dummy", mockConsumer, mockWriter)

	msg := consumerMessage()
	mockPartitionConsumerClient.On("Messages").Return(msg)
	mockPartitionConsumerClient.On("Errors").Return(nil)

	mockConsumer.On("Partitions", mock.Anything).Return([]int32{1}, nil)
	mockConsumer.On("ConsumePartition", mock.Anything, mock.Anything, mock.Anything).Return(mockPartitionConsumerClient, nil)
	mockConsumer.On("Close").Return(nil)
	mockWriter.On("InsertPatientProfile", mock.Anything, mock.Anything).Return(fmt.Errorf("insert failed")).Once()

	err := c.ConsumeMessage(context.Background())
	assert.NoError(t, err)

}

func consumerMessage() <-chan *sarama.ConsumerMessage {
	// Create a regular, two-way channel.
	c := make(chan *sarama.ConsumerMessage, 1)

	go func() {
		defer close(c)

		// Do stuff
		c <- &sarama.ConsumerMessage{
			Headers:        nil,
			Timestamp:      time.Time{},
			BlockTimestamp: time.Time{},
			Key:            []byte(`1`),
			Value:          []byte(`{ "PatientId": "1", "FirstName": "John", "LastName": "Doe", "IsPregnant": "False", "Sex": "Female"}`),
			Topic:          "patient-profile",
			Partition:      1,
			Offset:         0,
		}
	}()

	// Returning it, implicitly converts it to read-only,
	// as per the function return type.
	return c
}
