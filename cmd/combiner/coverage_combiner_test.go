package main

import (
	"os"
	"sort"
	"strings"
	"testing"
)

// Helper function to create a temporary file with given content
func createTempFile(t *testing.T, prefix string, content string) string {
	tmpFile, err := os.CreateTemp("", prefix)
	if err != nil {
		t.Fatalf("Failed to create temp file: %s", err)
	}
	if content != "" {
		_, err = tmpFile.WriteString(content)
		if err != nil {
			t.Fatalf("Failed to write to temp file: %s", err)
		}
	}
	tmpFile.Close()
	return tmpFile.Name()
}

// Helper function to validate the output against expected output
func validateOutput(t *testing.T, outputFile string, expectedOutput string) {
	// Read the output file
	outputBytes, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %s", err)
	}

	// Split and sort output and expected output for comparison
	output := strings.TrimSpace(string(outputBytes))
	outputLines := strings.Split(output, "\n")
	expectedLines := strings.Split(strings.TrimSpace(expectedOutput), "\n")

	sort.Strings(outputLines)
	sort.Strings(expectedLines)

	// Compare line counts
	if len(outputLines) != len(expectedLines) {
		t.Fatalf("Output does not match expected result. Number of lines differ.\nGot:\n%s\nExpected:\n%s",
			strings.Join(outputLines, "\n"), strings.Join(expectedLines, "\n"))
	}

	// Compare content line by line
	for i := range outputLines {
		if outputLines[i] != expectedLines[i] {
			t.Fatalf("Output mismatch at line %d.\nGot:\n%s\nExpected:\n%s",
				i+1, strings.Join(outputLines, "\n"), strings.Join(expectedLines, "\n"))
		}
	}
}

// Test case: Valid input produces correct output
func TestValidInputTest(t *testing.T) {
	testInput := `mode: set
github.com/example/file.go:1.1,2.1 1 0
github.com/example/file.go:1.1,2.1 1 1
github.com/example/file.go:2.1,3.1 2 0
github.com/example/file.go:3.1,4.1 1 0`

	expectedOutput := `mode: set
github.com/example/file.go:1.1,2.1 1 1
github.com/example/file.go:2.1,3.1 2 0
github.com/example/file.go:3.1,4.1 1 0`

	inputFile := createTempFile(t, "coverage-input", testInput)
	defer os.Remove(inputFile)

	outputFile := createTempFile(t, "coverage-output", "")
	defer os.Remove(outputFile)

	// Execute the main function
	os.Args = []string{"clean_coverage", inputFile, outputFile}
	main()

	// Validate the output
	validateOutput(t, outputFile, expectedOutput)
}

// Test case: Malformed input lines
func TestMalformedInputTest(t *testing.T) {
	testInput := `mode: set
malformed_input_line`
	expectedOutput := `mode: set`

	inputFile := createTempFile(t, "coverage-input", testInput)
	defer os.Remove(inputFile)

	outputFile := createTempFile(t, "coverage-output", "")
	defer os.Remove(outputFile)

	// Execute the main function
	os.Args = []string{"clean_coverage", inputFile, outputFile}
	main()

	// Validate the output
	validateOutput(t, outputFile, expectedOutput)
}
