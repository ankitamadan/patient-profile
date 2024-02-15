package consumer

import (
	"github.com/Shopify/sarama"

	"heidi_health/internal/store"
)

type Consumer struct {
	topicName      string
	writeStore     store.StoreWriter
	consumerClient ConsumerClient
}

func NewConsumer(topicName string, consumerClient ConsumerClient, writeStore store.StoreWriter) *Consumer {
	return &Consumer{topicName: topicName, writeStore: writeStore, consumerClient: consumerClient}
}

//go:generate mockery --name ConsumerClient
type ConsumerClient interface {
	Partitions(topic string) ([]int32, error)
	ConsumePartition(topic string, partition int32, offset int64) (sarama.PartitionConsumer, error)
	Close() error
}

//go:generate mockery --name PartitionConsumerClient
type PartitionConsumerClient interface {
	Messages() <-chan *sarama.ConsumerMessage
	Errors() <-chan *sarama.ConsumerError
	AsyncClose()
	Close() error
	HighWaterMarkOffset() int64
}
