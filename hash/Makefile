# Makefile

# Variables
BINARY_NAME=hash
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
	go mod init github.com/username/hash
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
	./$(BINARY) --input "Hello, World!"
	./$(BINARY) --input "Hello, World!" --algorithm sha1
	./$(BINARY) --input "Hello, World!" --algorithm sha512
	./$(BINARY) --input "Hello, World!" --algorithm md5

# Phony targets
.PHONY: setup build run clean complete
