# Makefile

# Variables
BINARY_NAME=smtpclient
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
SERVER_NAME=$(SERVER)
PORT_NUMBER=$(PORT)
FROM_ADDR=$(FROM)
TO_ADDR=$(TO)
SUBJ=$(SUBJECT)
BODYDATA=$(BODY)

# Setup dependencies
setup:
	go mod init github.com/username/smtpclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(SERVER_NAME) $(PORT_NUMBER) "$(FROM_ADDR)" "$(TO_ADDR)" "$(SUBJ)" "$(BODYDATA)"

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) 127.0.0.1 23456 "tess.t@example.com" "person@example.com" "\"Test email\"" "\"This is a test email\""

# Phony targets
.PHONY: setup build run clean complete
