# Makefile

# Variables
BINARY_NAME=telnetclient
ADDR=$(ADDRESS)

# Setup dependencies
setup:
	go mod init github.com/username/telnetclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(ADDR)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) 127.0.0.1:3456

# Phony targets
.PHONY: setup build run clean complete
