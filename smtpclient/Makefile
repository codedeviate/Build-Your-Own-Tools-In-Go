# Makefile

# Variables
BINARY_NAME=smtpclient
SERVER_NAME=$(SERVER)
PORT_NUMBER=$(PORT)
FROM_ADDR=$(FROM)
TO_ADDR=$(TO)
SUBJ=$(SUBJECT)
BODYDATA=$(BODY)

# Setup dependencies
setup:
	go mod init github.com/username/smtpclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(SERVER_NAME) $(PORT_NUMBER) "$(FROM_ADDR)" "$(TO_ADDR)" "$(SUBJ)" "$(BODYDATA)"

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) 127.0.0.1 23456 "tess.t@example.com" "person@example.com" "\"Test email\"" "\"This is a test email\""

# Phony targets
.PHONY: setup build run clean complete
