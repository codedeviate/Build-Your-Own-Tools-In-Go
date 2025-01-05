package compression

import (
	"compress/flate"
	"fmt"
	"io"
	"os"
)

func HandleDeflate(action, file string) {
	switch action {
	case "compress":
		compressDeflate(file)
	case "decompress":
		decompressDeflate(file)
	default:
		fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
	}
}

func compressDeflate(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file + ".deflate")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer, err := flate.NewWriter(outFile, flate.DefaultCompression)
	if err != nil {
		fmt.Println("Error creating deflate writer:", err)
		return
	}
	defer writer.Close()

	_, err = io.Copy(writer, inFile)
	if err != nil {
		fmt.Println("Error compressing file:", err)
	}
	fmt.Println("File compressed to", file+".deflate")
}

func decompressDeflate(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file[:len(file)-8]) // Remove ".deflate"
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	reader := flate.NewReader(inFile)
	defer reader.Close()

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Error decompressing file:", err)
	}
	fmt.Println("File decompressed to", outFile.Name())
}
