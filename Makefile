build:
	@go build -o bin/go-reservation cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/go-reservation

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# todo: remove the connection string and place it in ENV
migrate-up:
	@migrate -path db/migration/ -database "postgresql://postgres:gobank@localhost:5432/postgres?sslmode=disable" -verbose up

# todo: remove the connection string and place it in ENV
migrate-down:
	@migrate -path db/migration/ -database "postgresql://postgres:gobank@localhost:5432/postgres?sslmode=disable" -verbose down
