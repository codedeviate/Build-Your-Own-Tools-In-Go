package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/username/hash/lib"
)

func main() {
	algorithm := flag.String("algorithm", "sha256", "Hash algorithm: md5, sha1, sha256, sha512")
	input := flag.String("input", "", "Input string")
	filePath := flag.String("file", "", "Path to input file")
	flag.Parse()

	var data string
	if *filePath != "" {
		// Read from file
		result, err := lib.HashFile(*filePath, *algorithm)
		if err != nil {
			fmt.Printf("Error hashing file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(result)
		return
	} else if *input == "" {
		// Check if input is coming from a pipe
		fi, err := os.Stdin.Stat()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if fi.Mode()&os.ModeNamedPipe != 0 {
			// Read from stdin
			reader := bufio.NewReader(os.Stdin)
			inputBytes, err := io.ReadAll(reader)
			if err != nil {
				fmt.Printf("Error reading from stdin: %v\n", err)
				os.Exit(1)
			}
			data = string(inputBytes)
		} else {
			fmt.Println("Enter input string:")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				data = scanner.Text()
			}
		}
	} else {
		data = *input
	}

	result, err := lib.HashData(data, *algorithm)
	if err != nil {
		fmt.Printf("Error hashing data: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
