FROM golang:1.25.4-alpine

WORKDIR /app

RUN apk add --no-cache bash

COPY ./ ./

RUN go mod download

RUN chmod +x ./tester/entrypoint.sh

ENTRYPOINT [ "./tester/entrypoint.sh" ]