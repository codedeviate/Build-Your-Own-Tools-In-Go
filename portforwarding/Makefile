# Makefile

# Variables
BINARY_NAME=portforwarding
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
LISTEN=$(LISTEN_PORT)
SEND=$(TARGET)

# Setup dependencies
setup:
	go mod init github.com/username/portforwarding
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(LISTEN) $(SEND)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 3526 www.google.com:80

# Phony targets
.PHONY: setup build run clean complete
