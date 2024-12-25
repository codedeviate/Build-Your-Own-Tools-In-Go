# Makefile

# Variables
BINARY_NAME=portscanner
PROTOCOL=$(PROTOCOL)
HOST=$(HOST)
START_PORT=$(START_PORT)
END_PORT=$(END_PORT)

# Setup dependencies
setup:
	go mod init github.com/username/portscanner
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(PROTOCOL) $(HOST) $(START_PORT) $(END_PORT)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) tcp scanme.nmap.org 1 1024

# Phony targets
.PHONY: setup build run clean complete
