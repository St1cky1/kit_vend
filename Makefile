.PHONY: help proto gen-proto build run dev clean test

PROTOC := /usr/local/Cellar/protobuf/33.1/bin/protoc

help:
	@echo "Available commands:"
	@echo "  make proto       - Generate Go code from proto files"
	@echo "  make build       - Build the server"
	@echo "  make run         - Run the server"
	@echo "  make dev         - Run server in development mode"
	@echo "  make test        - Run tests"
	@echo "  make clean       - Clean build artifacts"

proto: gen-proto

gen-proto:
	@echo "Generating gRPC code from proto files..."
	$(PROTOC) -I. -I$(shell go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=paths=source_relative:. \
		--go-grpc_out=paths=source_relative:. \
		--grpc-gateway_out=paths=source_relative:. \
		pb/v1/vending_machine.proto
	@echo "Proto code generated successfully!"

build:
	@echo "Building server..."
	go build -o bin/server ./cmd/server/
	@echo "Build complete: bin/server"

run: build
	@echo "Running server..."
	./bin/server

proto-gen-doc:
	@echo "Generating API documentation from proto files..."
	mkdir -p docs
	$(PROTOC) -I. -I$(shell go env GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--doc_out=./docs \
		--doc_opt=markdown,API.md \
		pb/v1/vending_machine.proto
	@echo "API documentation generated: docs/API.md"

dev:
	@echo "Running server in development mode..."
	go run ./cmd/server/main.go

test:
	@echo "Running tests..."
	go test -v ./...

clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean
	@echo "Clean complete!"

mod-tidy:
	@echo "Tidying dependencies..."
	go mod tidy
	@echo "Dependencies tidied!"

lint:
	@echo "Running linter..."
	golangci-lint run ./...

fmt:
	@echo "Formatting code..."
	gofmt -s -w .
	@echo "Code formatted!"
