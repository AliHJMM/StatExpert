package main

import (
	"fmt"
	"log"
	"os"
	"STATEXPERT/functions"
)

func main(){
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0])
	}

	filename := os.Args[1]

	array, err := functions.ReadNumbersFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading from file: %s", err)
	}

	if len(array) == 0 {
		log.Println("No valid numbers to process.")
		return
	}

	average, median, variance, stddev := functions.CalculateStatistics(array)

	functions.PrintStatistics(average, median, variance, stddev)
}