# Makefile

# Variables
BINARY_NAME=smtpserver
PORT=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/smtpserver
	go get github.com/emersion/go-sasl
	go get github.com/emersion/go-smtp
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(PORT)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
