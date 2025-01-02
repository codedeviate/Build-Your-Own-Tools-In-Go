# Makefile

# Variables
BINARY_NAME=gomake
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif
BINARY=$(BINARY_NAME)$(EXT)
STR=$(HELLO)

# Setup dependencies
setup:
	go mod init github.com/username/go-make
	go get github.com/spf13/cobra@latest

build:
	go build -o $(BINARY)

test: build
	./$(BINARY) echo HELLO=$(STR)

echo:
	@echo Hello, $(STR)!

complete: clean setup build
	make test HELLO=world

clean:
	$(RM) $(BINARY)
	$(RM) go.mod
	$(RM) go.sum
