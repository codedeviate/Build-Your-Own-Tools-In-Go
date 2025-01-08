# Makefile

# Variables
BINARY_NAME=loadbalancer
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)

# Setup dependencies
setup:
	go mod init github.com/username/loadbalancer
	go get github.com/gorilla/mux
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY)

# Phony targets
.PHONY: setup build run clean complete
