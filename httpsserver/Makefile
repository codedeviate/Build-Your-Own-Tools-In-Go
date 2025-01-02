# Makefile

# Variables
BINARY_NAME=httpsserver
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
PORT=$(PORT)

# Setup dependencies
setup:
	go mod init github.com/username/httpsserver
	go mod tidy
	openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes -subj "/C=SE/ST=Skåne/L=Malmö/O=Example/OU=Development/CN=example.com"

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(PORT) cert.pem key.pem

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum
	$(RM) cert.pem
	$(RM) key.pem

complete: clean setup build
	./$(BINARY) 6789 cert.pem key.pem

# Phony targets
.PHONY: setup build run clean complete
