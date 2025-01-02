# Makefile

# Variables
BINARY_NAME=watcher
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
CONFIG_FILE=$(CONFIG)

# Setup dependencies
setup:
	go mod init github.com/username/watcher
	go get github.com/fsnotify/fsnotify
	go get github.com/gobwas/glob
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(CONFIG_FILE) $(PATH)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) config.json ../dummydata

# Phony targets
.PHONY: setup build run clean complete
