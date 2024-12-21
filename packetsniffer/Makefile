# Makefile

# Variables
BINARY_NAME=packetsniffer
INTERFACE=$(INTERFACE)

# Setup dependencies
setup:
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	sudo ./$(BINARY_NAME) $(INTERFACE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)

# Phony targets
.PHONY: setup build run clean