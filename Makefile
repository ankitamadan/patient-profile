mod-tidy:
	go mod tidy
.PHONY: mod-tidy

mod-dl:
	go mod download
.PHONY: mod-dl

## run unit tests
test:
	go test ./... -covermode=count -coverprofile=coverage.out -coverpkg=./... -json > tests.out
.PHONY: test

## build: build the producer
compile-producer:
	CGO_ENABLED=0 go build -o ./out/producer ./cmd/producer
.PHONY: compile-producer

## build: build the consumer
compile-consumer:
	CGO_ENABLED=0 go build -o ./out/consumer ./cmd/consumer
.PHONY: compile-consumer

run-producer: compile-producer
	./out/producer
.PHONY: run-producer

run-consumer: compile-consumer
	./out/consumer
.PHONY: run-consumer

local-dev-stack:
	docker-compose -f localstack-compose.yml  down && docker-compose -f localstack-compose.yml build  --no-cache && docker compose -f localstack-compose.yml up
