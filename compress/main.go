package main

import (
	"fmt"
	"os"

	"github.com/username/compress/compression"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run main.go <compress|decompress> <format> <file>")
		return
	}

	action := os.Args[1]
	format := os.Args[2]
	file := os.Args[3]

	switch format {
	case "gzip":
		compression.HandleGzip(action, file)
	case "zlib":
		compression.HandleZlib(action, file)
	case "deflate":
		compression.HandleDeflate(action, file)
	case "lzw":
		compression.HandleLzw(action, file)
	case "snappy":
		compression.HandleSnappy(action, file)
	default:
		fmt.Println("Unsupported format. Supported formats: gzip, zlib, deflate, lzw, snappy")
	}
}
