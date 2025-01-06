# Makefile

# Variables
BINARY_NAME=hexdump
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
FILENAME=$(FILE)

# Setup dependencies
setup:
	go mod init github.com/username/hexdump
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) file=$(FILENAME)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) -file ./$(BINARY)

# Phony targets
.PHONY: setup build run clean complete
