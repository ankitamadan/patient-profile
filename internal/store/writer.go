package store

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type WriteStore struct {
	client              DynamoDBWriter
	patientProfileTable string
}

func NewWriteStore(client DynamoDBWriter, patientProfileTable string) WriteStore {
	return WriteStore{client, patientProfileTable}
}

func (w WriteStore) InsertPatientProfile(ctx context.Context, patientProfile PatientProfile) error {
	patientProfile.UpdatedAt = time.Now()
	patientProfileUpdated, err := attributevalue.MarshalMap(patientProfile)
	if err != nil {
		return err
	}

	_, err = w.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(w.patientProfileTable),
		Item:      patientProfileUpdated,
	})

	return err
}

//go:generate mockery --name DynamoDBWriter
type DynamoDBWriter interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}
