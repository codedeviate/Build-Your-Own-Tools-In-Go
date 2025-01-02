# Makefile

# Variables
BINARY_NAME=portscanner
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
PROTOCOL=$(PROTOCOL)
HOST=$(HOST)
START_PORT=$(START_PORT)
END_PORT=$(END_PORT)

# Setup dependencies
setup:
	go mod init github.com/username/portscanner
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(PROTOCOL) $(HOST) $(START_PORT) $(END_PORT)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) tcp scanme.nmap.org 1 1024

# Phony targets
.PHONY: setup build run clean complete
