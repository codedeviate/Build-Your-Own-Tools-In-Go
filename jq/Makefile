# Makefile

# Variables
BINARY_NAME=jq
INTERFACE=$(INTERFACE)

# Setup dependencies
setup:
	go mod init github.com/username/jq
	go get github.com/tidwall/gjson
	go get github.com/clbanning/mxj
	go get github.com/jszwec/csvutil
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

complete: clean setup build
	./$(BINARY_NAME) en0

# Phony targets
.PHONY: setup build run clean complete
