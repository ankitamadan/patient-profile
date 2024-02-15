package main

import (
	"log"
	"net/http"

	"github.com/Shopify/sarama"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"heidi_health/internal/api"
	"heidi_health/internal/config"
)

func main() {

	envConfig := config.NewEnvironmentConfig()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	brokersUrl := []string{envConfig.KafkaHost}
	producerClient, err := connectProducer(brokersUrl)
	if err != nil {
		log.Fatal("cannot connect to kafka client")
	}

	a := api.NewAPI(envConfig, producerClient)

	router.Post("/patientprofile", a.PublishPatientProfile())

	if err = http.ListenAndServe(":3000", router); err != nil {
		log.Fatal("cannot start the server")
	}

}

func connectProducer(brokersUrl []string) (api.ProducerClient, error) {

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokersUrl, config)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
