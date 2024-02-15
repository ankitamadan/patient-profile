package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"

	"heidi_health/internal/store"
)

func (c Consumer) ConsumeMessage(ctx context.Context) error {

	partitions, err := c.consumerClient.Partitions(c.topicName)
	if err != nil {
		panic(err)
	}

	consumer, err := c.consumerClient.ConsumePartition(c.topicName, partitions[0], sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	fmt.Println("Consumer started ")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// Count how many message processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount++
				var patientProfile store.PatientProfile
				if msg != nil && msg.Value != nil {
					if err = json.Unmarshal(msg.Value, &patientProfile); err != nil {
						doneCh <- struct{}{}
					}
					if err = c.writeStore.InsertPatientProfile(ctx, patientProfile); err != nil {
						doneCh <- struct{}{}
					}
					fmt.Printf("Received message Count %d: | Topic(%s) | PatientId(%s) | FirstName(%s) | LastName(%s) \n", msgCount, msg.Topic, patientProfile.PatientID, patientProfile.FirstName, patientProfile.LastName)
				}
			case <-ctx.Done():
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			case <-sigchan:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")

	if err = c.consumerClient.Close(); err != nil {
		return err
	}

	return nil
}
