# Makefile

# Variables
BINARY_NAME=portknocking
HOST=$(HOSTNAME)
DELAY=$(DELAYMS)
PORT1=$(PORT_1)
PORT2=$(PORT_2)
# Setup dependencies
setup:
	go mod init github.com/username/portknocking
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(HOST) $(DELAY) $(PORT1) $(PORT2)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
