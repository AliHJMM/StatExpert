package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	filename := os.Args[1]

	array, err := ReadNumbersFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading from file: %s", err)
	}

	if len(array) == 0 {
		log.Println("No valid numbers to process.")
		return
	}
}