package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "<file_path>")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	header := make([]byte, 0x100)
	_, err = file.Read(header)
	if err != nil {
		fmt.Printf("Error reading file header: %v\n", err)
		return
	}

	fmt.Println("File Type:", identifyFileType(header))
}

func identifyFileType(header []byte) string {
	switch {
	case string(header[:4]) == "\x7fELF":
		return identifyELF(header)
	case string(header[:2]) == "MZ":
		return identifyDOS(header)
	case string(header[:4]) == "PE\x00\x00":
		return "Portable Executable (PE)"
	case string(header[:4]) == "\xca\xfe\xba\xbe":
		return "Mach-O 32-bit"
	case string(header[:4]) == "\xbe\xba\xfe\xca":
		return "Mach-O 32-bit"
	case string(header[:4]) == "\xfe\xed\xfa\xcf":
		return "Mach-O 64-bit"
	case string(header[:4]) == "\xcf\xfa\xed\xfe":
		return "Mach-O 64-bit"
	default:
		fmt.Printf("%X %X %X %X", header[0], header[1], header[2], header[3])
		return "Unknown format"
	}
}

func identifyELF(header []byte) string {
	if header[4] == 1 {
		return "ELF 32-bit"
	} else if header[4] == 2 {
		return "ELF 64-bit"
	}
	return "Unknown ELF type"
}

func identifyDOS(header []byte) string {
	peOffset := int(header[0x3c])
	if string(header[peOffset:peOffset+4]) == "PE\x00\x00" {
		archType := string(header[peOffset+4 : peOffset+6])
		if archType == "\x64\x86" {
			return "Portable Executable (PE) 64-bit"
		} else if archType == "\x4c\x01" {
			return "Portable Executable (PE) 32-bit"
		}
		optionalHeaderSize := int(header[peOffset+20])
		if optionalHeaderSize == 224 {
			return "Portable Executable (PE) 32-bit"
		} else if optionalHeaderSize == 240 {
			return "Portable Executable (PE) 64-bit"
		}
		return "Portable Executable (PE)"
	}
	return "DOS Executable (MZ)"
}
