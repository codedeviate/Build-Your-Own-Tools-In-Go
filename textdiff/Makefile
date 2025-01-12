# Makefile

# Variables
BINARY_NAME=textdiff
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
FILE_1=$(OLD_FILE)
FILE_2=$(NEW_FILE)

# Setup dependencies
setup:
	go mod init github.com/username/textdiff
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(FILE_1) $(FILE_2)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) ../dummydata/file1.json ../dummydata/file2.json

# Phony targets
.PHONY: setup build run clean complete
