# Makefile

# Variables
BINARY_NAME=encodedecode
INPUT_DATA=$(INPUT)

# Setup dependencies
setup:
	go mod init github.com/username/encodedecode
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	sudo ./$(BINARY_NAME) --input "$(INPUT_DATA)"

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) --input "Hello, World!" | ./$(BINARY_NAME) --mode decode

# Phony targets
.PHONY: setup build run clean complete
