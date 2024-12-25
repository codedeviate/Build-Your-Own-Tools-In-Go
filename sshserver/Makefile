# Makefile

# Variables
BINARY_NAME=sshserver
PORT=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/sshserver
	go get golang.org/x/crypto/ssh
	go mod tidy
	openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(PORT) key.pem

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum
	rm -f cert.pem
	rm -f key.pem

complete: clean setup build
	./$(BINARY_NAME) 4567 key.pem

# Phony targets
.PHONY: setup build run clean complete
