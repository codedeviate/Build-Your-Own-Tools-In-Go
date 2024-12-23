# Makefile

# Variables
BINARY_NAME=webscraper
URL=$(URL)
SELECTOR=$(SELECTOR)

# Setup dependencies
setup:
	go mod init github.com/username/webscraper
	go get github.com/PuerkitoBio/goquery
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(URL) $(SELECTOR)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
