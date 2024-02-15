package api_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"heidi_health/internal/api"
	"heidi_health/internal/api/mocks"
	envConfig "heidi_health/internal/config"
)

func TestAPI_PushPatientProfileToQueue_Success(t *testing.T) {
	mockProducer := mocks.NewProducerClient(t)
	a := api.NewAPI(envConfig.NewEnvironmentConfig(), mockProducer)

	mockProducer.On("SendMessage", mock.Anything).Return(int32(1), int64(1), nil)

	err := a.PushPatientProfileToQueue([]byte(`{"value":{"type":"JSON","data":{"name":"test"}}}`))
	assert.NoError(t, err)
}

func TestAPI_PushPatientProfileToQueue_Fail(t *testing.T) {
	mockProducer := mocks.NewProducerClient(t)
	a := api.NewAPI(envConfig.NewEnvironmentConfig(), mockProducer)

	mockProducer.On("SendMessage", mock.Anything).Return(int32(1), int64(1), fmt.Errorf("failed to produce message"))

	err := a.PushPatientProfileToQueue([]byte(`{"value":{"type":"JSON","data":{"name":"test"}}}`))
	assert.Error(t, err)
}
