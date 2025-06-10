.PHONY: test lint build clean fmt

# Build the CLI binary
build:
	go build -o base58 ./cmd/base58

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

# Run benchmarks
bench:
	go test -bench=. -benchmem

# Run benchmarks multiple times for stability
bench-stability:
	go test -bench=. -count=3

# Run benchmarks and save to file
bench-save:
	go test -bench=. -benchmem > benchmark_results.txt

# Run optimization comparison benchmarks
bench-compare:
	go test -bench=BenchmarkCompare -benchmem

# Run optimized version benchmarks
bench-optimized:
	go test -bench=BenchmarkEncodeOptimized -benchmem
	go test -bench=BenchmarkDecodeOptimized -benchmem

# Default target
all: check build