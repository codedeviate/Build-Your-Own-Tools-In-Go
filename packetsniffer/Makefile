# Makefile

# Variables
BINARY_NAME=packetsniffer
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
	go build -o $(BINARY_NAME) main.go

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
	sudo ./$(BINARY_NAME) $(INTERFACE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
ifeq ($(OS),Darwin)
	# macOS: Default interface is en0
	./$(BINARY_NAME) en0
endif
ifeq ($(OS),Linux)
	# Linux: Default interface is eth0
	./$(BINARY_NAME) eth0
endif

# Phony targets
.PHONY: setup build run clean complete
