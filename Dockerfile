# syntax=docker/dockerfile:1

## Build
FROM golang:1.19.2 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ENV GO119MODULE=on 
ENV GOOS=linux 
ENV GOARCH=amd64 
ENV CGO_ENABLED=0

RUN go mod tidy 
RUN go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bot .

# cert stage
FROM alpine:latest as certs
RUN apk --message add ca-certificates


## Deploy
FROM scratch

LABEL maintainer="SoXX <soxx@fenpa.ws>"

WORKDIR /app

COPY --from=build /app/bot .
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt


ENTRYPOINT [ "/app/bot" ]