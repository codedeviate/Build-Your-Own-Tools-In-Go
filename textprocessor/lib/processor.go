// processor.go
package lib

import (
	"bufio"
	"os"
	"strings"
)

// ProcessText processes the input text and applies transformations.
func ProcessText(text string) string {
    // Example transformation: convert text to uppercase
    return strings.ToUpper(text)
}

// ReadFile reads the content of a file.
func ReadFile(filePath string) (string, error) {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    return string(data), nil
}

// ReadStdin reads the content from standard input.
func ReadStdin() (string, error) {
    reader := bufio.NewReader(os.Stdin)
    var sb strings.Builder
    for {
        input, err := reader.ReadString('\n')
        if err != nil {
            break
        }
        sb.WriteString(input)
    }
    return sb.String(), nil
}
