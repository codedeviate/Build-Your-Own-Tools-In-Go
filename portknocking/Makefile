# Makefile

# Variables
BINARY_NAME=portknocking
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
HOST=$(HOSTNAME)
DELAY=$(DELAYMS)
PORT1=$(PORT_1)
PORT2=$(PORT_2)
# Setup dependencies
setup:
	go mod init github.com/username/portknocking
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(HOST) $(DELAY) $(PORT1) $(PORT2)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 127.0.0.1 1000 1234 5678

# Phony targets
.PHONY: setup build run clean complete
