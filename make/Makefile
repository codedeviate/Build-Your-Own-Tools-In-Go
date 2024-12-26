setup:
	go mod init github.com/username/go-make
	go get github.com/spf13/cobra@latest

build:
	go build -o gomake

test: build
	./gomake echo

echo:
	@echo "Hello, World!"

clean:
	rm -f gomake
	rm -f go.mod
	rm -f go.sum
