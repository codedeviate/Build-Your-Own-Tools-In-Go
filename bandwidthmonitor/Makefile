# Makefile

# Variables
BINARY_NAME=bandwidthmonitor
INTERFACE=$(INTERFACE)

# Setup dependencies
setup:
	go mod init github.com/username/bandwidthmonitor
	go get github.com/google/gopacket
	go get github.com/google/gopacket/pcap
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	sudo ./$(BINARY_NAME) $(INTERFACE)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
