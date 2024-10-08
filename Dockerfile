FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY ./backend .

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
