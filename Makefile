SHELL:=/bin/bash -O extglob
BINARY=ms-api
VERSION=0.1.0

LDFLAGS=-ldflags "-X main.Version=${VERSION}"

build:
	go build ${LDFLAGS} -o ${BINARY} cmd/main.go

run:
	@go run cmd/main.go