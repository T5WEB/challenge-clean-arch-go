init:
	docker compose up -d
	docker compose exec mysql bash -c 'sleep 5 && mysql -uroot -proot -e "CREATE DATABASE IF NOT EXISTS orders;"'

down:
	docker compose down

start:
	docker compose start

stop:
	docker compose stop

createmigration:
	migrate create -ext=sql -dir=internal/infra/database/sql/migrations -seq init

migrate:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

migratedown:
	migrate -path=internal/infra/database/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down

.PHONY: migrate migratedown createmigration stop init down start