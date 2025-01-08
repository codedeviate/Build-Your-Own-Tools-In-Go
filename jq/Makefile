# Makefile

# Variables
BINARY_NAME=jq
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
S_FILENAME=$(FILENAME)
S_QUERY=$(QUERY)
S_FORMAT=$(FORMAT)

# Setup dependencies
setup:
	go mod init github.com/username/jq
	go get github.com/tidwall/gjson
	go get github.com/clbanning/mxj
	go get github.com/jszwec/csvutil
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) --file $(S_FILENAME) --query $(S_QUERY) --format $(S_FORMAT)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) --file ../dummydata/file1.json --query "key2" --format "json"

# Phony targets
.PHONY: setup build run clean complete
