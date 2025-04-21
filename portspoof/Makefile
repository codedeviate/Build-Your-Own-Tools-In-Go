# Makefile

# Variables
BINARY_NAME=portspoof
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)

# Setup dependencies
setup:
	go mod init portspoof
	go mod tidy
	openssl req -new -x509 -days 365 -nodes -out cert.pem -keyout key.pem

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum
	$(RM) certs/cert.pem
	$(RM) certs/key.pem

complete: clean setup build
	./$(BINARY)

# Phony targets
.PHONY: setup build run clean complete
