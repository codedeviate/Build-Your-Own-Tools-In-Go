# Makefile

# Variables
BINARY_NAME=webscraper
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
URL=$(URL)
SELECTOR=$(SELECTOR)

# Setup dependencies
setup:
	go mod init github.com/username/webscraper
	go get github.com/PuerkitoBio/goquery
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	./$(BINARY) $(URL) $(SELECTOR)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum

complete: clean setup build
	./$(BINARY) https://github.com/codedeviate/Build-Your-Own-Tools-In-Go/blob/master/README.md article

# Phony targets
.PHONY: setup build run clean complete
