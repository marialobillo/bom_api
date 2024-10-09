# Define the binary name
BINARY_NAME=api

# Build the application
build:
	@go build -o bin/$(BINARY_NAME) cmd/main.go

# Run the application
run: build
	@./bin/$(BINARY_NAME)

# Test the application
test:
	@go test -v ./...

# Clean up the generated files
clean:
	@rm -rf bin/$(BINARY_NAME)
