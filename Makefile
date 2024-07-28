# Define variables
BINARY_NAME=global-entry-slot-notifier
SRC_DIR=./cmd/main.go
WITH_FLAGS=

# Default target
.PHONY: all
all: deps build

# Build the binary
.PHONY: build
build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(SRC_DIR)

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Clean build files
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)


# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod tidy

# Display help
.PHONY: help
help:
	@echo "Makefile targets:"
	@echo "  build     - Build the binary"
	@echo "  test      - Run tests"
	@echo "  clean     - Clean build files"
	@echo "  deps      - Install dependencies"
	@echo "  help      - Display this help message"
