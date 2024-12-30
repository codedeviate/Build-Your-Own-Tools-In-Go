# Makefile

# Variables
BINARY_NAME=hash
INPUT_DATA=$(INPUT)

# Setup dependencies
setup:
	go mod init github.com/username/hash
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
	./$(BINARY_NAME) --input "Hello, World!"
	./$(BINARY_NAME) --input "Hello, World!" --algorithm sha1
	./$(BINARY_NAME) --input "Hello, World!" --algorithm sha512
	./$(BINARY_NAME) --input "Hello, World!" --algorithm md5

# Phony targets
.PHONY: setup build run clean complete
