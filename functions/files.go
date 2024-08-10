package functions

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadNumbersFromFile(filename string) ([]int, error) {
	var array []int
	readingFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer readingFile.Close()

	file := bufio.NewScanner(readingFile)

	for file.Scan() {
		line := strings.TrimSpace(file.Text())
		if line == "" {
			continue 
		}
		if !isValidNumber(line) {
			log.Printf("Invalid number detected: '%s'", line)
			continue 
		}
		parsedValue, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Error parsing line '%s': %s", line, err)
			continue
		}
		array = append(array, parsedValue)
	}

	return array, nil
}
