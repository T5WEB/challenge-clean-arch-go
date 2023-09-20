# challenge-clean-arch-go

Fazer a listagem das orders com endpoint rest[x], query graphQL [x], endpoint gRPC[x].

## Commands

### To up or down containers docker

##### make init

##### make down

### To up or stop containers docker again

##### make start

##### make stop

### To create or to erase migrations

##### make migrate or make migratedown

### To run all services

##### go run main.go wire_gen.go

### To execute MySQL

##### docker compose exec mysql bash

##### after execute: mysql -p

### To generate folder pb of protocol buffer and grpc

##### protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

### To access the client EVANS of the gRPC

##### evans -r repl

    - Type: package pb
    - Type: service order_service
    - Type: call CreateOrder or ListOrders

Filling inputs of according

### To acces RabbitMQ

##### To access localhost:15672

admin: guest
pass: guest
