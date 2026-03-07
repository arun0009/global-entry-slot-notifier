# Define variables
BINARY_NAME=global-entry-slot-notifier
SRC_DIR=./cmd/main.go

.PHONY: help all build test clean deps
.DEFAULT_GOAL := help

help:
	@echo "Makefile targets:"
	@echo "  make       - Show this help"
	@echo "  make all   - Install deps and build"
	@echo "  make build - Build the binary"
	@echo "  make test  - Run tests"
	@echo "  make clean - Remove binary"
	@echo "  make deps  - go mod tidy"

all: deps build

build:
	@echo "Building $(BINARY_NAME)..."
	@go build -o $(BINARY_NAME) $(SRC_DIR)

test:
	@echo "Running tests..."
	@go test -v ./...

clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)

deps:
	@echo "Installing dependencies..."
	@go mod tidy
