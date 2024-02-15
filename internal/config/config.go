package config

import (
	"log"

	"github.com/caarlos0/env/v6"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type EnvConfig struct {
	ServerPort int    `env:"SERVER_PORT" envDefault:"3000"`
	TopicName  string `env:"TOPIC_NAME" envDefault:"patient.profile.updated"`
	DdbHost    string `env:"DDB_HOST" envDefault:"http://localhost:4566"`
	DDBName    string `env:"DDB_NAME" envDefault:"patient-profile"`
	AWSRegion  string `env:"AWS_REGION" envDefault:"ap-southeast-2"`
	KafkaHost  string `env:"KAFKA_HOST" envDefault:"localhost:9094"`
}

func (c EnvConfig) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ServerPort, validation.Required),
		validation.Field(&c.TopicName, validation.Required),
		validation.Field(&c.KafkaHost, validation.Required),
		validation.Field(&c.DdbHost, validation.Required),
		validation.Field(&c.DDBName, validation.Required),
		validation.Field(&c.AWSRegion, validation.Required),
	)
}

func NewEnvironmentConfig() *EnvConfig {
	cfg := &EnvConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("cannot find configs for server")
	}

	errors := cfg.Validate()
	if errors != nil {
		log.Fatal("Configuration validation failed with errors")
	}

	return cfg
}
