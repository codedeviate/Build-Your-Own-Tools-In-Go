# Makefile

# Variables
BINARY_NAME=telnetserver
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
PORT=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/telnetserver
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(PORT)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 3456

# Phony targets
.PHONY: setup build run clean complete
