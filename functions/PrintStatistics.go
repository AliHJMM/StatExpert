package functions

import "fmt"

func PrintStatistics(average, median, variance, stddev float64) {
	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", stddev)
}
