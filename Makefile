.PHONY: build, run

build:
	@go mod download
	@mkdir -p bin
	@go build -o bin/ ./...

run:
	@./bin/rest-api-with-new-routing