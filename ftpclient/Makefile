# Makefile

# Variables
BINARY_NAME=ftpclient
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
SERVER_ADDRESS=$(SERVER)
USERNAME=$(USER)
PASSW=$(PASSWORD)
FILE=$(FILEPATH)

# Setup dependencies
setup:
	go mod init github.com/username/ftpclient
	go get github.com/jlaffaye/ftp
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(SERVER_ADDRESS) $(USERNAME) $(PASSW) $(FILE)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 127.0.0.1:3253 admin admin /file1.txt

# Phony targets
.PHONY: setup build run clean complete
