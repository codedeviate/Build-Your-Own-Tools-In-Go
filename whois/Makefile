# Makefile

# Variables
BINARY_NAME=whois
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
DOMAIN=$(DOMAIN)

# Setup dependencies
setup:
	go mod init github.com/username/whois
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(DOMAIN)

# Clean the build
clean:
	$(RM) $(BINARY)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) codedv8.dev

# Phony targets
.PHONY: setup build run clean complete
