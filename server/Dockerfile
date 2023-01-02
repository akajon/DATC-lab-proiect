FROM golang:1.19-alpine3.17 AS base

WORKDIR /app

COPY . .

WORKDIR /app/cmd/monolith

RUN go build -o main

EXPOSE 8081

CMD [ "./main" ]