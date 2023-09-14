# challenge-clean-arch-go

## Commands

#### To up containers docker

### docker compose up -d

##### To run all services

### go run main.go wire_gen.go

#### To execute MySQL

### docker compose exec mysql bash

#### To access the client EVANS of the gRPC

### evans -r repl

    - Type: package pb
    - Type: service order_service
    - Type: call CreateOrder

Filling inputs of according

#### To acces RabbitMQ

### To access localhost:15672

admin: guest
pass: guest
