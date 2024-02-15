package api_test

import (
	"bytes"

	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"heidi_health/internal/api"
	"heidi_health/internal/api/mocks"
	envConfig "heidi_health/internal/config"
)

func TestAPI_Handler_400(t *testing.T) {
	mockProducer := mocks.NewProducerClient(t)
	a := api.NewAPI(envConfig.NewEnvironmentConfig(), mockProducer)
	handler := a.PublishPatientProfile()

	req := httptest.NewRequest("POST", "/api/v1", nil)
	rw := httptest.NewRecorder()

	handler.ServeHTTP(rw, req)

	assert.Equal(t, 400, rw.Code)
}

func TestAPI_Handler_200(t *testing.T) {
	mockProducer := mocks.NewProducerClient(t)

	a := api.NewAPI(envConfig.NewEnvironmentConfig(), mockProducer)
	mockProducer.On("SendMessage", mock.Anything).Return(int32(1), int64(1), nil)

	handler := a.PublishPatientProfile()

	req := httptest.NewRequest("POST", "/api/v1", bytes.NewReader([]byte(`{ "PatientId": "2", "FirstName": "John", "LastName": "Doe", "IsPregnant": "False"}`)))
	rw := httptest.NewRecorder()

	handler.ServeHTTP(rw, req)

	assert.Equal(t, 201, rw.Code)
}
