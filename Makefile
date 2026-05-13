.PHONY: build test vet lint tidy cover

build:
	go build ./...

test:
	go test ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

tidy:
	go mod tidy

cover:
	go test -coverprofile=coverage.out ./...
