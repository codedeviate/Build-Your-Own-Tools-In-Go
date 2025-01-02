# Makefile

# Variables
BINARY_NAME=httpclient
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
URL=$(URL)

# Setup dependencies
setup:
	go mod init github.com/username/httpclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(URL)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) http://www.google.com

# Phony targets
.PHONY: setup build run clean complete
