.PHONY: build test clean run help

all: test build

build:
	@echo "Building Flip 7 Simulator..."
	@go build -o flip7-simulator ./cmd/
	@echo "✅ Build complete!"

test:
	@echo "Running tests..."
	@go test -v ./internal/game/
	@echo "✅ All tests passed!"

clean:
	@rm -f flip7-simulator
	@echo "✅ Clean complete!"

run: build
	@./flip7-simulator

quick: build
	@./flip7-simulator -games 100

help:
	@echo "Available targets:"
	@echo "  build  - Build the simulator"
	@echo "  test   - Run unit tests"
	@echo "  clean  - Remove build artifacts"
	@echo "  run    - Build and run with default settings"
	@echo "  quick  - Build and run with 100 games"
	@echo "  help   - Show this help message"
