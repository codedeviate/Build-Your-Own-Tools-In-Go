# Makefile

# Variables
BINARY_NAME=netscan
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
HOST=$(HOST)
START_PORT=$(START_PORT)
END_PORT=$(END_PORT)

# Setup dependencies
setup:
	go mod init github.com/username/netscan
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(HOST) $(START_PORT) $(END_PORT)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 127.0.0.1 1024 2048

# Phony targets
.PHONY: setup build run clean complete
