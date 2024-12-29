// main.go
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/username/jq/lib"
)

func main() {
	file := flag.String("file", "", "Path to the file (CSV, JSON, or XML)")
	query := flag.String("query", "", "Query to apply to the data")
	format := flag.String("format", "json", "Format of the input data (csv, json, or xml)")
	flag.Parse()

	if *query == "" {
		fmt.Println("Usage: jq_program --file <file_path> --query <query> --format <csv|json|xml>")
		os.Exit(1)
	}

	var data string
	var err error

	if *file != "" {
		data, err = lib.ReadFile(*file)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			os.Exit(1)
		}
	} else {
		data, err = lib.ReadStdin()
		if err != nil {
			fmt.Printf("Error reading stdin: %v\n", err)
			os.Exit(1)
		}
	}

	var result string
	switch *format {
	case "json":
		result = lib.ProcessJSON(data, *query)
	case "xml":
		result = lib.ProcessXML(data, *query)
	case "csv":
		result = lib.ProcessCSV(data, *query)
	default:
		fmt.Println("Unsupported format. Use 'csv', 'json', or 'xml'.")
		os.Exit(1)
	}

	fmt.Println(result)
}
