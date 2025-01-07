// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/username/textprocessor/lib"
)

func main() {
	file := flag.String("file", "", "Path to the file")
	flag.Parse()

	var data string
	var err error

	if *file != "" {
		data, err = lib.ReadFile(*file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		data, err = lib.ReadStdin()
		if err != nil {
			fmt.Printf("Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	}

	result := lib.ProcessText(data)
	fmt.Println(result)
}
