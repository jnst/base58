# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a high-performance Go Base58 encoding/decoding library with CLI tool, implementing the Bitcoin-standard character set. The project emphasizes performance optimization, test-driven development, and comprehensive benchmarking.

## Core Development Commands

### Building
```bash
make build          # Build CLI binary
make clean          # Clean build artifacts
```

### Testing Strategy
```bash
# Fast testing for modified code only (speed prioritized)
go test -run=TestSpecificFunction
go test ./cmd -run=TestCLI         # Test CLI functionality only

# Full test suite (use sparingly)
make test           # All tests
go test -v          # Library tests with verbose output
```

### Code Quality (Required workflow)
```bash
make check          # Format + lint + test (complete workflow)
make fmt            # Format code only
make lint           # Lint only
```

### Benchmarking
```bash
make bench                  # All benchmarks
make bench-stability        # Run benchmarks multiple times
make bench-save             # Save results to file
```

## Architecture

### Core Implementation
- **base58.go**: Core encode/decode functions using big.Int arithmetic with object pool optimization
- **cmd/base58/main.go**: CLI application with file/stdin support

The implementation uses object pools for big.Int and strings.Builder reuse to achieve high performance.

### Performance Architecture
- Object pools for big.Int and strings.Builder reuse
- Optimized buffer size calculations
- Memory allocation reduction (up to 99.9% reduction achieved)
- Bitcoin-standard Base58 character set: `123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz`

## Development Workflow

**Critical**: Every implementation must follow this complete workflow:
1. **Implementation** - Write/modify code
2. **Lint + Test** - `make check` must pass
3. **Documentation** - Update README.md, benchmarks, etc.

This is a complete set - do not skip any step.

## Testing Guidelines

### Speed-Focused Testing
- Run only tests for modified code to maintain development speed
- Use specific test patterns: `go test -run=TestSpecificFunction`
- Avoid full test suite during development iterations

### Test Coverage
- **base58_test.go**: Core functionality tests with edge cases
- **cmd/base58/main_test.go**: CLI functionality tests
- Fuzz testing for random data validation
- Round-trip testing for correctness

## Configuration

### Go Version
- Requires Go 1.18 or later (for `any` type in object pools)

### Linting
- golangci-lint with strict configuration (.golangci.yml)
- Test files have relaxed rules for benchmark code
- Errcheck exclusions for benchmark functions

## Performance Focus

This is a performance-critical library:
- Memory allocation reduction is primary goal
- Benchmark-driven development approach
- Comprehensive benchmarking to measure performance improvements
- Object pool patterns for resource reuse

## Key Files

- **Makefile**: All development commands
- **.golangci.yml**: Strict linting configuration  
- **base58_bench_test.go**: Comprehensive benchmarking
- **cmd/base58/main.go**: CLI application
- **cmd/base58/main_test.go**: CLI functionality tests
- **BENCHMARK.md**: Performance analysis documentation
- **OPTIMIZATION_RESULTS.md**: Optimization achievements documentation

## Common Tasks

### Adding New Functions
1. Implement in base58.go
2. Write tests in base58_test.go
3. Add benchmarks in base58_bench_test.go
4. Run `make check`
5. Update documentation

### Performance Optimization
1. Profile with benchmarks first
2. Implement optimization
3. Validate correctness with existing tests
4. Measure improvement with benchmarks
5. Update optimization documentation

### CLI Modifications
1. Modify cmd/base58/main.go
2. Test with cmd/base58/main_test.go
3. Validate with actual CLI usage
4. Update help text and examples