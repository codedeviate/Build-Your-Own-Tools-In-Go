# Makefile

# Variables
BINARY_NAME=loganalyzer
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
FILENAME=$(FILE)
LOG_TYPE=$(TYPE)

# Setup dependencies
setup:
	go mod init github.com/username/loganalyzer
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) --file $(FILENAME) --type $(LOG_TYPE)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) --file ../dummydata/apache_access.log --type apache_access

# Phony targets
.PHONY: setup build run clean complete
