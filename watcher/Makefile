# Makefile

# Variables
BINARY_NAME=watcher
CONFIG_FILE=$(CONFIG)

# Setup dependencies
setup:
	go mod init github.com/username/watcher
	go get github.com/fsnotify/fsnotify
	go get github.com/gobwas/glob
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(CONFIG_FILE) $(PATH)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

complete: clean setup build
	./$(BINARY_NAME) config.json ../dummydata

# Phony targets
.PHONY: setup build run clean complete
