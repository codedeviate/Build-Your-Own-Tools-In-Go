# Makefile

# Variables
BINARY_NAME=traceroute
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
HOST=$(HOST)

# Setup dependencies
setup:
	go mod init github.com/username/traceroute
	go get "golang.org/x/net/icmp"
	go get "golang.org/x/net/ipv4"
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(HOST)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) www.google.com

# Phony targets
.PHONY: setup build run clean complete
