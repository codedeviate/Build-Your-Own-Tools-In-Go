# Makefile

# Variables
BINARY_NAME=ipgeolocation
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
IP=$(IP)

# Setup dependencies
setup:
	go mod init github.com/username/ipgeolocation
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(IP)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 8.8.8.8

# Phony targets
.PHONY: setup build run clean complete
