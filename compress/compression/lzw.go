package compression

import (
	"compress/lzw"
	"fmt"
	"io"
	"os"
)

func HandleLzw(action, file string) {
	switch action {
	case "compress":
		compressLzw(file)
	case "decompress":
		decompressLzw(file)
	default:
		fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
	}
}

func compressLzw(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file + ".lzw")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Use LSB (least significant bit) order
	writer := lzw.NewWriter(outFile, lzw.LSB, 8)
	defer writer.Close()

	_, err = io.Copy(writer, inFile)
	if err != nil {
		fmt.Println("Error compressing file:", err)
	}
	fmt.Println("File compressed to", file+".lzw")
}

func decompressLzw(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file[:len(file)-4]) // Remove ".lzw"
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	// Use LSB (least significant bit) order
	reader := lzw.NewReader(inFile, lzw.LSB, 8)
	defer reader.Close()

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Error decompressing file:", err)
	}
	fmt.Println("File decompressed to", outFile.Name())
}
