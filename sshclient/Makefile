# Makefile

# Variables
BINARY_NAME=sshclient
USERNAME=$(USER)
PASSW=$(PASSWORD)
HOSTNAME=$(HOST)
PORT_NUMBER=$(PORT)
CMD=$(COMMAND)

# Setup dependencies
setup:
	go mod init github.com/username/sshclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(USERNAME) $(PASSW) $(HOSTNAME) $(PORT_NUMBER) "$(CMD)"

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
