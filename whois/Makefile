# Makefile

# Variables
BINARY_NAME=whois
DOMAIN=$(DOMAIN)

# Setup dependencies
setup:
	go mod init github.com/username/whois
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(DOMAIN)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
