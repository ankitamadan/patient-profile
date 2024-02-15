package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"heidi_health/internal/config"
)

func Test_envConfig_Validate(t *testing.T) {
	validConfig := config.EnvConfig{
		ServerPort: 123,
		TopicName:  "dummy",
		KafkaHost:  "host",
		DDBName:    "dummytable",
		AWSRegion:  "region",
		DdbHost:    "host",
	}

	err := validConfig.Validate()
	assert.NoError(t, err)
}
