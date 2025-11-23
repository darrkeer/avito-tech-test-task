FROM golang:1.25.4-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o server ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY ./migrations ./migrations
CMD ["./server"]
