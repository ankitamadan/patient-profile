package api

import (
	"time"

	"github.com/Shopify/sarama"

	"heidi_health/internal/config"
)

type API struct {
	config         *config.EnvConfig
	producerClient ProducerClient
}

func NewAPI(config *config.EnvConfig, producerClient ProducerClient) *API {
	return &API{config: config, producerClient: producerClient}
}

type PatientProfilePayload struct {
	PatientID  string    `json:"PatientId"`
	FirstName  string    `json:"FirstName"`
	LastName   string    `json:"LastName"`
	IsPregnant string    `json:"IsPregnant"`
	Sex        string    `json:"Sex"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}

//go:generate mockery --name PublishPatientProfile
type PublishPatientProfile interface {
	PushPatientProfileToQueue(message []byte) error
}

//go:generate mockery --name ProducerClient
type ProducerClient interface {
	SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error)
}
