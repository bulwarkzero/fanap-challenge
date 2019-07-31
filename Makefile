default: build

run:
	@go run main.go
.PHONY: run

build:
	@go build
.PHONY: build
