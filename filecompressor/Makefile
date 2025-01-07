# Makefile

# Variables
BINARY_NAME=filecompressor
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
S_ACTION=$(ACTION)
S_SOURCE=$(SOURCE)
S_TARGET=$(TARGET)

# Setup dependencies
setup:
	go mod init github.com/username/filecompressor
	go mod tidy

# Build the binary
build:
	go build -o $(BINARY) main.go

# Run the application
run: build
	sudo ./$(BINARY) --$(S_ACTION) --source $(S_SOURCE) --target $(S_TARGET)

# Clean the build
clean:
	$(RM) $(BINARY_NAME)
	$(RM) go.mod
	$(RM) go.sum
	$(RM) main.go.gz

complete: clean setup build
	./$(BINARY) --compress --source ./main.go --target ./main.go.gz

# Phony targets
.PHONY: setup build run clean complete
