# Makefile

# Variables
BINARY_NAME=portforwarding
LISTEN=$(LISTEN_PORT)
SEND=$(TARGET)

# Setup dependencies
setup:
	go mod init github.com/username/portforwarding
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(LISTEN) $(SEND)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) 3526 www.google.com

# Phony targets
.PHONY: setup build run clean complete
