# Makefile

# Variables
BINARY_NAME=packetsniffer
SERVER_ADDRESS=$(SERVER)
USERNAME=$(USER)
PASSW=$(PASSWORD)
FILE=$(FILEPATH)

# Setup dependencies
setup:
	go mod init github.com/username/packetsniffer
	go get github.com/jlaffaye/ftp
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(SERVER_ADDRESS) $(USERNAME) $(PASSW) $(FILE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) 127.0.0.1:3254 admin admin /file1.txt

# Phony targets
.PHONY: setup build run clean complete
