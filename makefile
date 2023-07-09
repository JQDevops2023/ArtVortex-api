.PHONY: clean

BINARY_NAME=server

all: run

build:
	go build cmd/server/${BINARY_NAME}.go

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root artvortex 

dropdb:
	docker exec -it postgres12 dropdb artvortex

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nft2023 -d postgres:12-alpine
run:
	./${BINARY_NAME}

start:
	make build
	make run

generate:
	go get github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate

clean:
	go clean -cache
	rm -f ${BINARY_NAME}
dep:
	go mod download

lint:
	golangci-lint run --enable-all