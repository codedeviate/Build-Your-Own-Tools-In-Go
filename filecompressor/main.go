// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/username/filecompressor/lib"
)

func main() {
	compress := flag.Bool("compress", false, "Compress the file")
	decompress := flag.Bool("decompress", false, "Decompress the file")
	source := flag.String("source", "", "Source file path")
	target := flag.String("target", "", "Target file path")
	flag.Parse()

	if *source == "" || *target == "" {
		fmt.Println("Usage: file_compressor --compress|--decompress --source <source_file> --target <target_file>")
		os.Exit(1)
	}

	var err error
	if *compress {
		err = lib.CompressFile(*source, *target)
	} else if *decompress {
		err = lib.DecompressFile(*source, *target)
	} else {
		fmt.Println("Specify either --compress or --decompress")
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Operation failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Operation successful")
}
