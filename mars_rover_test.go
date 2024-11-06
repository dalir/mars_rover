package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Helper function to run the binary with given input and capture output
func runMarsRover(inputFile string) (string, error) {
	cmd := exec.Command("./mars_rover", "-i", inputFile)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func TestMarsRover(t *testing.T) {
	// Define directories
	inputDir := "./input"
	expectedDir := "./output"

	// List all test files in input directory
	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		t.Fatalf("Failed to read input directory: %v", err)
	}

	for _, file := range files {
		inputFile := filepath.Join(inputDir, file.Name())
		expectedFile := filepath.Join(expectedDir, file.Name())

		t.Run(file.Name(), func(t *testing.T) {
			// Run the Mars Rover program with the input file
			output, err := runMarsRover(inputFile)
			if err != nil {
				t.Fatalf("Failed to run mars_rover on %s: %v", inputFile, err)
			}

			// Read expected output
			expectedOutput, err := ioutil.ReadFile(expectedFile)
			if err != nil {
				t.Fatalf("Failed to read expected output file %s: %v", expectedFile, err)
			}

			// Compare the program output with expected output
			if strings.TrimSpace(output) != strings.TrimSpace(string(expectedOutput)) {
				t.Errorf("Test %s failed:\nExpected:\n%s\nGot:\n%s", file.Name(), expectedOutput, output)
			}
		})
	}
}
