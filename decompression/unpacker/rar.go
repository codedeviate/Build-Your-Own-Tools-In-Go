// unpacker/rar.go
package unpacker

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/nwaples/rardecode"
)

// Unrar extracts a .rar file to the specified destination folder.
func Unrar(src string, dest string) error {
	// Open the rar file
	file, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open rar file: %v", err)
	}
	defer file.Close()

	// Create a new rar reader
	r, err := rardecode.NewReader(file, "")
	if err != nil {
		return fmt.Errorf("failed to create rar reader: %v", err)
	}

	// Extract files from the rar archive
	for {
		header, err := r.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read next file: %v", err)
		}

		// Extract each file
		err = extractRarFile(header, r, dest)
		if err != nil {
			return err
		}
	}
	return nil
}

// extractRarFile extracts a single file from the rar archive.
func extractRarFile(header *rardecode.FileHeader, r *rardecode.Reader, dest string) error {
	destPath := filepath.Join(dest, header.Name)

	// Create the destination file
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer destFile.Close()

	// Copy the content from the rar file to the destination file
	_, err = io.Copy(destFile, r)
	if err != nil {
		return fmt.Errorf("failed to copy file content: %v", err)
	}
	return nil
}
