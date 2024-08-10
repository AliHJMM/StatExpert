package functions

import (
	"math"
	"sort"
)

// CalculateStatistics computes the average, median, variance, and standard deviation of an integer slice.
func CalculateStatistics(array []int) (average, median, variance, stddev float64) {
	sort.Ints(array) // Sort the array to calculate median
	// Calculate the average
	total := 0.0
	for _, num := range array {
		total += float64(num)
	}

	average = math.Round(total / float64(len(array)))

	// Calculate the median
	if len(array)%2 == 0 {
		// Even number of elements: average of the two middle elements
		median = math.Ceil(float64((array[(len(array)/2)-1])+(array[(len(array)/2)])) / 2)
	} else {
		// Odd number of elements: middle element
		median = math.Ceil(float64(array[len(array)/2]))
	}

	// Calculate variance
	var varianceSum float64
	for _, num := range array {
		varianceSum += math.Pow((float64(num) - average), 2)
	}
	variance = math.Round(varianceSum / float64(len(array)))
	// Calculate standard deviation
	stddev = math.Round(math.Sqrt(variance))

	return average, median, variance, stddev
}
