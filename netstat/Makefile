# Makefile

# Variables
BINARY_NAME=netstat

# Setup dependencies
setup:
	go mod init github.com/username/netstat
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME)

# Phony targets
.PHONY: setup build run clean complete
