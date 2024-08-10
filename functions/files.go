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
}
