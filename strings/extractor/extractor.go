// extractor/extractor.go
package extractor

import (
	"os"

	"github.com/username/strings/utils"
)

// Minimum length of printable strings to extract
const MinStringLength = 4

// ExtractStrings reads a file and returns a slice of printable strings.
func ExtractStrings(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result []string
	var currentString []byte

	buf := make([]byte, 4096) // Buffer size
	for {
		n, err := file.Read(buf)
		if err != nil && err.Error() != "EOF" {
			return nil, err
		}

		// Process bytes in the buffer
		for i := 0; i < n; i++ {
			c := buf[i]
			if utils.IsPrintable(c) {
				currentString = append(currentString, c)
			} else {
				if len(currentString) >= MinStringLength {
					result = append(result, string(currentString))
				}
				currentString = nil
			}
		}

		// Break if we reached EOF
		if err != nil {
			break
		}
	}

	// Add any remaining string after the loop
	if len(currentString) >= MinStringLength {
		result = append(result, string(currentString))
	}

	return result, nil
}
