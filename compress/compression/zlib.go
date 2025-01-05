package compression

import (
	"compress/zlib"
	"fmt"
	"io"
	"os"
)

func HandleZlib(action, file string) {
	switch action {
	case "compress":
		compressZlib(file)
	case "decompress":
		decompressZlib(file)
	default:
		fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
	}
}

func compressZlib(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file + ".zlib")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer := zlib.NewWriter(outFile)
	defer writer.Close()

	_, err = io.Copy(writer, inFile)
	if err != nil {
		fmt.Println("Error compressing file:", err)
	}
	fmt.Println("File compressed to", file+".zlib")
}

func decompressZlib(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file[:len(file)-5]) // Remove ".zlib"
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	reader, err := zlib.NewReader(inFile)
	if err != nil {
		fmt.Println("Error creating zlib reader:", err)
		return
	}
	defer reader.Close()

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Error decompressing file:", err)
	}
	fmt.Println("File decompressed to", outFile.Name())
}
