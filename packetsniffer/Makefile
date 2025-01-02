# Makefile

# Variables
BINARY_NAME=packetsniffer
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
INTERFACE=$(INTERFACE)
OS := $(shell uname)

# Setup dependencies
setup:
	go mod init github.com/username/packetsniffer
	go get github.com/google/gopacket
	go get github.com/google/gopacket/pcap
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
ifeq ($(OS),Darwin)
	# macOS: Default interface is en0
	INTERFACE ?= en0
endif
ifeq ($(OS),Linux)
	# Linux: Default interface is eth0, wlan0 could also be an option
	INTERFACE ?= eth0
endif
	sudo ./$(BINARY) $(INTERFACE)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
ifeq ($(OS),Darwin)
	# macOS: Default interface is en0
	./$(BINARY) en0
endif
ifeq ($(OS),Linux)
	# Linux: Default interface is eth0
	./$(BINARY) eth0
endif

# Phony targets
.PHONY: setup build run clean complete
