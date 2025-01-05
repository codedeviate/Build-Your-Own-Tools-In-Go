# Makefile

# Variables
BINARY_NAME=compress
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
S_ACTION=$(ACTION)
S_TYPE=$(TYPE)
S_FILE=$(FILE)

# Setup dependencies
setup:
	go mod init github.com/username/compress
	go get github.com/golang/snappy
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) $(S_ACTION) $(S_TYPE) $(S_INFILE)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) compress gzip main.go

# Phony targets
.PHONY: setup build run clean complete
