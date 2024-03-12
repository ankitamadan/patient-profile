package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/Shopify/sarama"

	envConfig "heidi_health/internal/config"
	c "heidi_health/internal/consumer"
	"heidi_health/internal/store"
)

func main() {

	cfg := envConfig.NewEnvironmentConfig()

	ctx := context.TODO()

	// Using the Config value, create the DynamoDB client
	ddbClient := createClient(ctx, cfg.AWSRegion, cfg.DdbHost)

	consumerClient, err := connectConsumer([]string{cfg.KafkaHost})
	if err != nil {
		panic(err)
	}

	writeStore := store.NewWriteStore(ddbClient, cfg.DDBName)

	consume := c.NewConsumer(cfg.TopicName, consumerClient, writeStore)

	if err = consume.ConsumeMessage(ctx); err != nil {
		panic(err)
	}

}

func createClient(ctx context.Context, awsRegion, awsEndpoint string) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(awsRegion),
		config.WithClientLogMode(aws.LogRequest|aws.LogRetries),
	)
	if err != nil {
		panic(err)
	}

	// nolint:staticcheck
	cfg.EndpointResolverWithOptions = aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           awsEndpoint,
			SigningRegion: awsRegion,
		}, nil
	})
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return dynamodb.NewFromConfig(cfg)
}

func connectConsumer(brokersUrl []string) (c.ConsumerClient, error) {
	cfg := sarama.NewConfig()
	cfg.Consumer.Return.Errors = true

	// Create new consumer
	conn, err := sarama.NewConsumer(brokersUrl, cfg)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
