.PHONY: all
all: execute

run:
	go run main.go
build:
	go build -o bin/main main.go
test:
	go test ./...
execute: build
	$(shell cd bin && ./main)

# TODO
# docker-build:
#            docker run -it -v $PWD:/search-api golang:1.14.7-alpine3.12

# go modules clean tiny
