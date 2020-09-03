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
image:
	docker build --tag search-api .
container:
	docker run -it -p 8000:8000 --name api-test search-api