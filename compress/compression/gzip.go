package compression

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func HandleGzip(action, file string) {
	switch action {
	case "compress":
		compressGzip(file)
	case "decompress":
		decompressGzip(file)
	default:
		fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
	}
}

func compressGzip(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file + ".gz")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer := gzip.NewWriter(outFile)
	defer writer.Close()

	_, err = io.Copy(writer, inFile)
	if err != nil {
		fmt.Println("Error compressing file:", err)
	}
	fmt.Println("File compressed to", file+".gz")
}

func decompressGzip(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file[:len(file)-3]) // Remove ".gz"
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	reader, err := gzip.NewReader(inFile)
	if err != nil {
		fmt.Println("Error creating gzip reader:", err)
		return
	}
	defer reader.Close()

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Error decompressing file:", err)
	}
	fmt.Println("File decompressed to", outFile.Name())
}
