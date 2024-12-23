# Makefile

# Variables
BINARY_NAME=webcrawler
URL=$(URL)

# Setup dependencies
setup:
	go mod init github.com/username/webcrawler
	go get github.com/PuerkitoBio/goquery
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY_NAME) main.go

# Run the application
run: build
	./$(BINARY_NAME) $(URL)

# Clean the build
clean:
	rm -f $(BINARY_NAME)
	rm -f go.mod
	rm -f go.sum

# Phony targets
.PHONY: setup build run clean
