package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run github.com/Arematics/go-tests/coverage_combiner.go <input_file> [<output_file>]")
		return
	}

	inputFile := os.Args[1]
	outputFile := "cleaned-coverage.out"
	if len(os.Args) > 2 {
		outputFile = os.Args[2]
	}

	coverageMap := make(map[string]string)

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		os.Exit(1) // Exit the program with error code 1
	}
	defer file.Close()

	scanner := scan(file, coverageMap)

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %s\n", err)

		os.Exit(1) // Exit the program with error code 1
	}

	// Write the cleaned coverage report
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)

		os.Exit(1) // Exit the program with error code 1
	}
	defer output.Close()

	// Write the mode header (if it exists first)
	if header, exists := coverageMap["header"]; exists {
		output.WriteString(header + "\n")
	}

	// Write the unique coverage lines
	for key, line := range coverageMap {
		if key == "header" {
			continue
		}
		output.WriteString(line + "\n")
	}

	fmt.Printf("Cleaned coverage file generated: %s\n", outputFile)
}

func scan(file *os.File, coverageMap map[string]string) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Preserve the mode header line
		if strings.HasPrefix(line, "mode:") {
			coverageMap["header"] = line
			continue
		}

		// Parse the line into fields
		parts := strings.Fields(line)
		if len(parts) < 3 {
			fmt.Printf("Invalid line format: %s\n", line)
			continue
		}

		key := parts[0] + ":" + parts[1] // Combine file name and range as unique key

		// Prioritize lines with a "1" hit count
		if existing, exists := coverageMap[key]; exists {
			if strings.HasSuffix(existing, "1") {
				continue // Keep the tested line
			}
		}

		coverageMap[key] = line // Add or replace the entry
	}
	return scanner
}
