# Makefile

# Variables
BINARY_NAME=traceroute
HOST=$(HOST)

# Setup dependencies
setup:
	go mod init github.com/username/traceroute
	go get "golang.org/x/net/icmp"
	go get "golang.org/x/net/ipv4"
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(HOST)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) www.google.com

# Phony targets
.PHONY: setup build run clean complete
