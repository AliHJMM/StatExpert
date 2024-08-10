package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadNumbersFromFile reads integers from a file and returns them as a slice of ints.
func ReadNumbersFromFile(filename string) ([]int, error) {
	var array []int

	// Open the specified file
	readingFile, err := os.Open(filename)
	if err != nil {
		return nil, err // Return error if the file cannot be opened
	}
	defer readingFile.Close() // Ensure the file is closed after reading

	// Create a scanner to read the file line by line
	file := bufio.NewScanner(readingFile)

	// Process each line in the file
	for file.Scan() {
		line := strings.TrimSpace(file.Text()) // Trim any whitespace from the line
		if line == "" {
			continue  // Skip empty lines
		}
		if !isValidNumber(line) {
			log.Printf("Invalid number detected: '%s'", line) // Log invalid numbers
			continue 
		}
		parsedValue, err := strconv.Atoi(line) // Convert the line to an integer
		if err != nil {
			log.Printf("Error parsing line '%s': %s", line, err) // Log parsing errors
			continue
		}
		array = append(array, parsedValue) // Add the valid integer to the array
	}

	return array, nil
}
