package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func (api *API) PublishPatientProfile() http.HandlerFunc {
	handler := func(rw http.ResponseWriter, r *http.Request) error {
		reqBody, _ := ioutil.ReadAll(r.Body)

		// Instantiate new Message struct
		patientProfile := new(PatientProfilePayload)
		patientProfile.UpdatedAt = time.Now()
		if err := json.Unmarshal(reqBody, &patientProfile); err != nil {
			err = WriteJson(rw, http.StatusBadRequest, "Error creating patient")
			return err
		}
		// convert body into bytes and send it to kafka
		patientProfileInBytes, err := json.Marshal(patientProfile)
		if err != nil {
			err = WriteJson(rw, http.StatusBadRequest, "Error creating patient")
			return err
		}
		err = api.PushPatientProfileToQueue(patientProfileInBytes)
		if err != nil {
			err = WriteJson(rw, http.StatusBadRequest, "Error publishing patient profile event")
			return err
		}

		// Return PatientProfile in JSON format
		return WriteJson(rw, http.StatusCreated, "patient created successfully")
	}

	return InternalServerErrorHandler(handler)
}

func WriteJson(rw http.ResponseWriter, statusCode int, payload interface{}) error {
	if payload == nil {
		rw.WriteHeader(statusCode)
		return nil
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	if err := json.NewEncoder(rw).Encode(payload); err != nil {
		return err
	}
	return nil
}

type HttpHandlerFunc func(http.ResponseWriter, *http.Request) error

func InternalServerErrorHandler(handler HttpHandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if err := handler(writer, request); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	}
}
