# Makefile

# Variables
BINARY_NAME=executable-inspector
EXT=
RM=rm -f
ifeq ($(OS),Windows_NT)
    EXT=.exe
	RM=del /Q
endif

# Setup dependencies
setup:
	go mod init github.com/username/executable-inspector
	go mod tidy

build:
	go build -o $(BINARY_NAME)$(EXT) main.go

build-windows-amd64: export GOOS=windows
build-windows-amd64: export GOARCH=amd64
build-windows-amd64:
	go build -o $(BINARY_NAME)_windows_amd64.exe main.go

build-linux-amd64: export GOOS=linux
build-linux-amd64: export GOARCH=amd64
build-linux-amd64:
	go build -o $(BINARY_NAME)_linux_amd64 main.go

build-darwin-amd64: export GOOS=darwin
build-darwin-amd64: export GOARCH=amd64
build-darwin-amd64:
	go build -o $(BINARY_NAME)_darwin_amd64 main.go

build-freebsd-amd64: export GOOS=freebsd
build-freebsd-amd64: export GOARCH=amd64
build-freebsd-amd64:
	go build -o $(BINARY_NAME)_freebsd_amd64 main.go

build-windows-386: export GOOS=windows
build-windows-386: export GOARCH=386
build-windows-386:
	go build -o $(BINARY_NAME)_windows_386.exe main.go

build-linux-386: export GOOS=linux
build-linux-386: export GOARCH=386
build-linux-386:
	go build -o $(BINARY_NAME)_linux_386 main.go

build-freebsd-386: export GOOS=freebsd
build-freebsd-386: export GOARCH=386
build-freebsd-386:
	go build -o $(BINARY_NAME)_freebsd_386 main.go

build-windows-arm: export GOOS=windows
build-windows-arm: export GOARCH=arm
build-windows-arm:
	go build -o $(BINARY_NAME)_windows_arm.exe main.go

build-linux-arm: export GOOS=linux
build-linux-arm: export GOARCH=arm
build-linux-arm:
	go build -o $(BINARY_NAME)_linux_arm main.go

build-freebsd-arm: export GOOS=freebsd
build-freebsd-arm: export GOARCH=arm
build-freebsd-arm:
	go build -o $(BINARY_NAME)_freebsd_arm main.go

build-windows-arm64: export GOOS=windows
build-windows-arm64: export GOARCH=arm64
build-windows-arm64:
	go build -o $(BINARY_NAME)_windows_arm64.exe main.go

build-linux-arm64: export GOOS=linux
build-linux-arm64: export GOARCH=arm64
build-linux-arm64:
	go build -o $(BINARY_NAME)_linux_arm64 main.go

build-darwin-arm64: export GOOS=darwin
build-darwin-arm64: export GOARCH=arm64
build-darwin-arm64:
	go build -o $(BINARY_NAME)_darwin_arm64 main.go

build-freebsd-arm64: export GOOS=freebsd
build-freebsd-arm64: export GOARCH=arm64
build-freebsd-arm64:
	go build -o $(BINARY_NAME)_freebsd_arm64 main.go

build-all: build build-windows-amd64 build-linux-amd64 build-darwin-amd64 build-freebsd-amd64 build-windows-386 build-linux-386 build-freebsd-386 build-windows-arm build-linux-arm build-freebsd-arm build-windows-arm64 build-linux-arm64 build-darwin-arm64 build-freebsd-arm64

run: build-all
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_windows_amd64.exe
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_linux_amd64
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_darwin_amd64
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_freebsd_amd64
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_windows_386.exe
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_linux_386
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_freebsd_386
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_windows_arm.exe
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_linux_arm
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_freebsd_arm
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_windows_arm64.exe
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_linux_arm64
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_darwin_arm64
	./$(BINARY_NAME)$(EXT) $(BINARY_NAME)_freebsd_arm64

clean:
	$(RM) $(BINARY_NAME)$(EXT)
	$(RM) $(BINARY_NAME)_windows_amd64.exe
	$(RM) $(BINARY_NAME)_linux_amd64
	$(RM) $(BINARY_NAME)_darwin_amd64
	$(RM) $(BINARY_NAME)_freebsd_amd64
	$(RM) $(BINARY_NAME)_windows_386.exe
	$(RM) $(BINARY_NAME)_linux_386
	$(RM) $(BINARY_NAME)_freebsd_386
	$(RM) $(BINARY_NAME)_windows_arm.exe
	$(RM) $(BINARY_NAME)_linux_arm
	$(RM) $(BINARY_NAME)_freebsd_arm
	$(RM) $(BINARY_NAME)_windows_arm64.exe
	$(RM) $(BINARY_NAME)_linux_arm64
	$(RM) $(BINARY_NAME)_darwin_arm64
	$(RM) $(BINARY_NAME)_freebsd_arm64
	$(RM) go.mod
	$(RM) go.sum
