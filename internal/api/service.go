package api

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"

	"heidi_health/internal/store"
)

func (a *API) PushPatientProfileToQueue(message []byte) error {

	var patientProfile store.PatientProfile
	_ = json.Unmarshal(message, &patientProfile)

	msg := &sarama.ProducerMessage{
		Topic: a.config.TopicName,
		// Kafka Message Keys
		// Alongside the message value, we can choose to send a message key and that key it could be a string/number(we are using PatientID in our case),
		// although if we don’t send the key, the key is set to null then the data will be sent in a Round Robin fashion. So that means the first message is going
		// to be sent to partition 0, and then your second message to partition 1 and then partition 2, and so on.
		// This is why it’s called Round Robin, but in case we send a key with our message, all the messages that share the same key will always go to the same partition.
		// For ordering, for a specific field, if we want all our events come in order, we need to make sure we have message key set as the unique identifier i.e patientID and so
		// we need to choose the message key equal to PatientID so that we have all the patients for that one specific patient ir order as pat of the same partition
		// Order is only guaranteed from within a partition.
		Key:   sarama.StringEncoder(patientProfile.PatientID),
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := a.producerClient.SendMessage(msg)
	if err != nil {
		return err
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", a.config.TopicName, partition, offset)

	return nil
}
