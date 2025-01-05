// utils/utils.go
package utils

import (
	"fmt"
	"strings"
)

// GetExtension checks the file extension of the given file.
func GetExtension(filePath string) string {
	tempFilePath := strings.ToLower(filePath)
	fileExtension := tempFilePath[strings.LastIndex(tempFilePath, ".")+1:]
	return fileExtension
}

// IsZip checks if the file is a .zip file.
func IsZip(filePath string) bool {
	fmt.Println(filePath)
	fmt.Println(GetExtension(filePath))
	return GetExtension(filePath) == "zip"
}

// IsRar checks if the file is a .rar file.
func IsRar(filePath string) bool {
	return GetExtension(filePath) == "rar"
}

// Is7z checks if the file is a .7z file.
func Is7z(filePath string) bool {
	return GetExtension(filePath) == "7z"
}
