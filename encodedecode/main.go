package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/username/encodedecode/lib"
)

func main() {
	mode := flag.String("mode", "encode", "Mode: encode or decode")
	method := flag.String("method", "base64", "Method: base64 or url")
	input := flag.String("input", "", "Input string")
	filePath := flag.String("file", "", "Path to input file")
	flag.Parse()

	var data string
	if *filePath != "" {
		// Read from file
		file, err := os.Open(*filePath)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		reader := bufio.NewReader(file)
		inputBytes, err := io.ReadAll(reader)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
		data = string(inputBytes)
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

	var result string
	var err error

	switch *method {
	case "base64":
		if *mode == "encode" {
			result = lib.Base64Encode(data)
		} else {
			result, err = lib.Base64Decode(data)
		}
	case "url":
		if *mode == "encode" {
			result = lib.URLEncode(data)
		} else {
			result, err = lib.URLDecode(data)
		}
	default:
		fmt.Println("Unsupported method. Use 'base64' or 'url'.")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(result)
}
