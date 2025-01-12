package main

import (
	"bufio"
	"fmt"
	"os"
)

// DiffResult represents the results of a diff operation.
type DiffResult struct {
	Added   []string // Lines that are in new text but not in old text
	Removed []string // Lines that are in old text but not in new text
}

// ReadFile reads the file content and returns it as a slice of strings (lines).
func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Diff compares two slices of strings (representing two versions of text).
// It returns a DiffResult containing lines added and removed.
func Diff(oldLines, newLines []string) DiffResult {
	var diff DiffResult

	oldSet := make(map[string]struct{})
	newSet := make(map[string]struct{})

	// Fill the maps with old and new lines
	for _, line := range oldLines {
		oldSet[line] = struct{}{}
	}
	for _, line := range newLines {
		newSet[line] = struct{}{}
	}

	// Find removed lines (present in old but not in new)
	for _, line := range oldLines {
		if _, exists := newSet[line]; !exists {
			diff.Removed = append(diff.Removed, line)
		}
	}

	// Find added lines (present in new but not in old)
	for _, line := range newLines {
		if _, exists := oldSet[line]; !exists {
			diff.Added = append(diff.Added, line)
		}
	}

	return diff
}

// PrintDiff prints the diff result to the console.
func PrintDiff(diff DiffResult) {
	fmt.Println("Removed lines:")
	for _, line := range diff.Removed {
		fmt.Printf("- %s\n", line)
	}

	fmt.Println("\nAdded lines:")
	for _, line := range diff.Added {
		fmt.Printf("+ %s\n", line)
	}
}

// main function to drive the diff comparison.
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: diff <old file> <new file>")
		os.Exit(1)
	}

	// Read the old and new files.
	oldFilePath := os.Args[1]
	newFilePath := os.Args[2]

	oldLines, err := ReadFile(oldFilePath)
	if err != nil {
		fmt.Printf("Error reading old file: %v\n", err)
		return
	}

	newLines, err := ReadFile(newFilePath)
	if err != nil {
		fmt.Printf("Error reading new file: %v\n", err)
		return
	}

	// Perform the diff comparison.
	diff := Diff(oldLines, newLines)

	// Print the results.
	PrintDiff(diff)
}
