package functions

import (
	"math"
	"sort"
)

func CalculateStatistics(array []int) (average, median, variance, stddev float64) {
	sort.Ints(array)
	total := 0.0
	for _, num := range array {
		total += float64(num)
	}

	average = math.Round(total / float64(len(array)))
}