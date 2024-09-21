# Stage 1: Build the Go binary
FROM golang:1.23-alpine AS builder

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to the workspace
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the rest of the application code to the container
COPY . .

# Build the Go application
RUN go build -o main ./cmd

# Stage 2: Create a lightweight image for running the Go binary
FROM alpine:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Expose the port on which the app will run
EXPOSE 8080

# Command to run the Go application
CMD ["./main"]
