# Makefile

# Variables
BINARY_NAME=httpserver
PORT=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/httpserver
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(PORT)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
