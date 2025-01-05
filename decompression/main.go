// main.go
package main

import (
	"fmt"
	"os"

	"github.com/username/decompression/unpacker"
	"github.com/username/decompression/utils"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <file-path> <destination-folder>")
		os.Exit(1)
	}

	filePath := os.Args[1]
	destFolder := os.Args[2]

	// Check the file type and call the appropriate unpacking function
	switch {
	case utils.IsZip(filePath):
		err := unpacker.Unzip(filePath, destFolder)
		if err != nil {
			fmt.Println("Error unpacking zip file:", err)
		} else {
			fmt.Println("Successfully unpacked zip file!")
		}
	case utils.IsRar(filePath):
		err := unpacker.Unrar(filePath, destFolder)
		if err != nil {
			fmt.Println("Error unpacking rar file:", err)
		} else {
			fmt.Println("Successfully unpacked rar file!")
		}
	case utils.Is7z(filePath):
		err := unpacker.Un7z(filePath, destFolder)
		if err != nil {
			fmt.Println("Error unpacking 7z file:", err)
		} else {
			fmt.Println("Successfully unpacked 7z file!")
		}
	default:
		fmt.Println("Unsupported file format:", filePath)
	}
}
