# Makefile

# Variables
BINARY_NAME=encodedecode
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
INPUT_DATA=$(INPUT)

# Setup dependencies
setup:
	go mod init github.com/username/encodedecode
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) --input "$(INPUT_DATA)"

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) --input "Hello, World!" | ./$(BINARY) --mode decode

# Phony targets
.PHONY: setup build run clean complete
