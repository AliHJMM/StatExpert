package main

import (
	"log"
	"os"
	"STATEXPERT/functions"
)

func main(){
	// Check if a filename argument is provided
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <filename>", os.Args[0]) // Print error and exit if no filename is provided
	}
	// Get the filename from command-line arguments
	filename := os.Args[1]
	// Read numbers from the specified file
	array, err := functions.ReadNumbersFromFile(filename)
	if err != nil {
		log.Fatalf("Error reading from file: %s", err) // Print error and exit if reading fails
	}
	// Check if the array is empty
	if len(array) == 0 {
		log.Println("No valid numbers to process.") // Log a message if there are no valid numbers
		return
	}
	// Calculate statistical metrics: average, median, variance, and standard deviation	
	average, median, variance, stddev := functions.CalculateStatistics(array)
	// Print the calculated statistics
	functions.PrintStatistics(average, median, variance, stddev)
}