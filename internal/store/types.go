package store

import (
	"context"
	"time"
)

type PatientProfile struct {
	PatientID  string    `dynamodbav:"PatientId"`
	FirstName  string    `dynamodbav:"FirstName"`
	Sex        string    `dynamodbav:"Sex"`
	LastName   string    `dynamodbav:"LastName"`
	IsPregnant string    `dynamodbav:"IsPregnant"`
	UpdatedAt  time.Time `dynamodbav:"UpdatedAt"`
}

//go:generate mockery --name StoreWriter
type StoreWriter interface {
	InsertPatientProfile(context.Context, PatientProfile) error
}
