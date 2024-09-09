build:
	@go build -o bin/go-reservation cmd/api/main.go

run: build
	@./bin/go-reservation

test:
	@go test -v ./...