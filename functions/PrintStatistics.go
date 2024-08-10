package functions

import "fmt"
// PrintStatistics outputs the average, median, variance, and standard deviation.
func PrintStatistics(average, median, variance, stddev float64) {
	fmt.Printf("Average: %.0f\n", average)  // Print the average
	fmt.Printf("Median: %.0f\n", median)  // Print the median
	fmt.Printf("Variance: %.0f\n", variance)  // Print the variance
	fmt.Printf("Standard Deviation: %.0f\n", stddev)  // Print the standard deviation
}
