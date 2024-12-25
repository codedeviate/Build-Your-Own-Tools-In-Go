# Makefile

# Variables
BINARY_NAME=netscan
HOST=$(HOST)
START_PORT=$(START_PORT)
END_PORT=$(END_PORT)

# Setup dependencies
setup:
	go mod init github.com/username/netscan
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(HOST) $(START_PORT) $(END_PORT)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) 127.0.0.1 1 1024

# Phony targets
.PHONY: setup build run clean complete
