# Makefile

# Variables
BINARY_NAME=taskscheduler
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)

# Setup dependencies
setup:
	go mod init github.com/username/taskscheduler
	go get github.com/robfig/cron/v3
	go mod tidy

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

complete: clean setup build
	./$(BINARY)

# Phony targets
.PHONY: setup build run clean complete
