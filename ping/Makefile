# Makefile

# Variables
BINARY_NAME=ping
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
TARGET=$(TARGET)

# Setup dependencies
setup:
	go mod init github.com/username/ping
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(TARGET)

# Clean the build
clean:
	$(RM) $(BINARY)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) www.google.com

# Phony targets
.PHONY: setup build run clean complete
