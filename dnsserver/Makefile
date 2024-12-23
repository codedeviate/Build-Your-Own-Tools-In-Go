# Makefile

# Variables
BINARY_NAME=dnsserver
PORT_NUMBER=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/dnsserver
	go get github.com/miekg/dns
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(PORT_NUMBER)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
