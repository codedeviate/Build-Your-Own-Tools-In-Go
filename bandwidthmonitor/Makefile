# Makefile

# Variables
BINARY_NAME=bandwidthmonitor
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
INTERFACE=$(INTERFACE)

# Setup dependencies
setup:
	go mod init github.com/username/bandwidthmonitor
	go get github.com/google/gopacket
	go get github.com/google/gopacket/pcap
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) $(INTERFACE)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	sudo ./$(BINARY) en0

# Phony targets
.PHONY: setup build run clean complete
