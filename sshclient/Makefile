# Makefile

# Variables
BINARY_NAME=sshclient
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
USERNAME=$(USER)
PASSW=$(PASSWORD)
HOSTNAME=$(HOST)
PORT_NUMBER=$(PORT)
CMD=$(COMMAND)

# Setup dependencies
setup:
	go mod init github.com/username/sshclient
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(USERNAME) $(PASSW) $(HOSTNAME) $(PORT_NUMBER) "$(CMD)"

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) "\"\"" "\"\"" "127.0.0.1" "4567" "ls -alF"

# Phony targets
.PHONY: setup build run clean complete
