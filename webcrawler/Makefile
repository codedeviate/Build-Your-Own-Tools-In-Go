# Makefile

# Variables
BINARY_NAME=webcrawler
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
URL=$(URL)

# Setup dependencies
setup:
	go mod init github.com/username/webcrawler
	go get github.com/PuerkitoBio/goquery
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(URL)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) https://github.com/codedeviate/Build-Your-Own-Tools-In-Go/blob/master/README.md

# Phony targets
.PHONY: setup build run clean complete
