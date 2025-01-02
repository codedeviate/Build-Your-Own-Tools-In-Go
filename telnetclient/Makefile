# Makefile

# Variables
BINARY_NAME=telnetclient
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
ADDR=$(ADDRESS)

# Setup dependencies
setup:
	go mod init github.com/username/telnetclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(ADDR)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 127.0.0.1:3456

# Phony targets
.PHONY: setup build run clean complete
