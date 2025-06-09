.PHONY: test lint build clean fmt

# Build the CLI binary
build:
	go build -o base58 ./cmd

# Run tests
test:
	go test -v ./...

# Run golangci-lint
lint:
	golangci-lint run

# Format code
fmt:
	gofmt -w .

# Clean build artifacts
clean:
	rm -f base58

# Run all checks (format, lint, test)
check: fmt lint test

# Install dependencies
deps:
	go mod download
	go mod tidy

# Install golangci-lint
install-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Full setup for development
setup: deps install-lint

# Default target
all: check build