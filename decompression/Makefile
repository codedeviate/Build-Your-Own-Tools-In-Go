# Makefile

# Variables
BINARY_NAME=decompression
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
S_FILE=$(FILE)
S_PATH=$(PATH)

# Setup dependencies
setup:
	go mod init github.com/username/decompression
	go get github.com/nwaples/rardecode
	go get github.com/mholt/archiver/v3
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) $(S_FILE) $(S_PATH)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	mkdir dummydata
	./$(BINARY) ../dummydata/test.zip ./dummydata

# Phony targets
.PHONY: setup build run clean complete
