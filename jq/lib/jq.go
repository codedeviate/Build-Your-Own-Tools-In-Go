// jq.go
package lib

import (
	"fmt"
	"io"
	"os"

	"github.com/clbanning/mxj"
	"github.com/jszwec/csvutil"
	"github.com/tidwall/gjson"
)

// ProcessJSON processes JSON input and applies a query.
func ProcessJSON(jsonData, query string) string {
	result := gjson.Get(jsonData, query)
	return result.String()
}

// ProcessXML processes XML input and applies a query.
func ProcessXML(xmlData, query string) string {
	mv, err := mxj.NewMapXml([]byte(xmlData))
	if err != nil {
		return fmt.Sprintf("Error parsing XML: %v", err)
	}

	result, err := mv.ValueForPath(query)
	if err != nil {
		return fmt.Sprintf("Error querying XML: %v", err)
	}

	return fmt.Sprintf("%v", result)
}

// ProcessCSV processes CSV input and applies a query.
func ProcessCSV(csvData, query string) string {
	var records []map[string]interface{}
	if err := csvutil.Unmarshal([]byte(csvData), &records); err != nil {
		return fmt.Sprintf("Error parsing CSV: %v", err)
	}

	for _, record := range records {
		if value, ok := record[query]; ok {
			return fmt.Sprintf("%v", value)
		}
	}
	return "Query not found"
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
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
