package store_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"heidi_health/internal/store"
	"heidi_health/internal/store/mocks"
)

func TestNewWriteStore(t *testing.T) {
	mockWriter := mocks.NewDynamoDBWriter(t)
	mockDbClient := store.NewWriteStore(mockWriter, "mock-table")
	assert.NotNil(t, mockDbClient)
}

func TestWriteStore_InsertPatientProfile_Success(t *testing.T) {
	mockWriter := new(mocks.DynamoDBWriter)

	patientProfile := store.PatientProfile{
		PatientID:  "1",
		FirstName:  "John",
		LastName:   "Dow",
		IsPregnant: "False",
		Sex:        "Female",
	}

	mockWriter.On("PutItem", mock.Anything, mock.Anything).Return(
		&dynamodb.PutItemOutput{}, nil)
	mockDbClient := store.NewWriteStore(mockWriter, "mock-table")

	err := mockDbClient.InsertPatientProfile(context.Background(), patientProfile)
	assert.NoError(t, err)
}

func TestWriteStore_InsertPatientProfile_Fail(t *testing.T) {
	mockWriter := new(mocks.DynamoDBWriter)

	patientProfile := store.PatientProfile{
		PatientID:  "1",
		FirstName:  "John",
		LastName:   "Dow",
		IsPregnant: "False",
		Sex:        "Female",
	}

	mockWriter.On("PutItem", mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("insertion failed"))
	mockDbClient := store.NewWriteStore(mockWriter, "mock-table")

	err := mockDbClient.InsertPatientProfile(context.Background(), patientProfile)
	assert.Error(t, err)
}
