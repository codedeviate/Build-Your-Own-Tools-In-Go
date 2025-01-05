package compression

import (
	"fmt"
	"io"
	"os"

	"github.com/golang/snappy"
)

func HandleSnappy(action, file string) {
	switch action {
	case "compress":
		compressSnappy(file)
	case "decompress":
		decompressSnappy(file)
	default:
		fmt.Println("Invalid action. Use 'compress' or 'decompress'.")
	}
}

func compressSnappy(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file + ".snappy")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	writer := snappy.NewBufferedWriter(outFile)
	defer writer.Close()

	_, err = io.Copy(writer, inFile)
	if err != nil {
		fmt.Println("Error compressing file:", err)
	}
	fmt.Println("File compressed to", file+".snappy")
}

func decompressSnappy(file string) {
	inFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer inFile.Close()

	outFile, err := os.Create(file[:len(file)-7]) // Remove ".snappy"
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outFile.Close()

	reader := snappy.NewReader(inFile)

	_, err = io.Copy(outFile, reader)
	if err != nil {
		fmt.Println("Error decompressing file:", err)
	}
	fmt.Println("File decompressed to", outFile.Name())
}
