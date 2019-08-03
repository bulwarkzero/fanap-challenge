default: build

run:
	@go run main.go
.PHONY: run

build: # @env CGO_ENABLED=0 GOOS=linux go build -a
	@go build
.PHONY: build
