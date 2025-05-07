# Simple Makefile for a Go project

# Build the application
all: build test

build:
	@echo "Building..."
	
	
	@go build -o breakout-web cmd/web/breakout-web.go

# Run the application
run:
	@go run cmd/web/breakout-web.go

# Test the application
test:
	@echo "Testing..."
	@go test ./... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f breakout-web

