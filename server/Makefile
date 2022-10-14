run:
	go run cmd/main.go

test: unit

unit:
	go test -v -cover -tags=unit ./...

# DB URL format
# dbdriver://username:password@host:port/dbname?param1=true&param2=false

DB_DRIVER=postgresql
USERNAME=program
PASSWORD=test
HOST=localhost
PORT=5432
DB_NAME=persons

DB_URL="$(DB_DRIVER)://$(USERNAME):$(PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable"

migrate_create:
	migrate create -ext sql -dir schema -seq init
	migrate create -ext sql -dir schema -seq seed

migrate_up:
	migrate -path schema -database $(DB_URL) -verbose up

migrate_down:
	migrate -path schema -database $(DB_URL) -verbose down
