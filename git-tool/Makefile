# Makefile

# Variables
BINARY_NAME=git-tool
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)

# Setup dependencies
setup:
	go mod init github.com/username/git-tool
	go get github.com/go-git/go-git/v5
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(ACTION)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) status

# Phony targets
.PHONY: setup build run clean complete
