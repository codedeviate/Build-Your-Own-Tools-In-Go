// unpacker/zip.go
package unpacker

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Unzip extracts a .zip file to the specified destination folder.
func Unzip(src string, dest string) error {
	// Open the zip file
	zipFile, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %v", err)
	}
	defer zipFile.Close()

	// Iterate over the files in the archive
	for _, file := range zipFile.File {
		err := extractZipFile(file, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

// extractZipFile extracts a single file from the zip archive.
func extractZipFile(file *zip.File, dest string) error {
	// Open the file inside the zip archive
	rc, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open file inside zip: %v", err)
	}
	defer rc.Close()

	// Prepare the destination file path
	destPath := filepath.Join(dest, file.Name)

	// Create the necessary directories if they don't exist
	if file.FileInfo().IsDir() {
		err = os.MkdirAll(destPath, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	} else {
		// Create the destination file
		destFile, err := os.Create(destPath)
		if err != nil {
			return fmt.Errorf("failed to create file: %v", err)
		}
		defer destFile.Close()

		// Copy the content from the zip file to the destination file
		_, err = io.Copy(destFile, rc)
		if err != nil {
			return fmt.Errorf("failed to copy file content: %v", err)
		}
	}
	return nil
}
