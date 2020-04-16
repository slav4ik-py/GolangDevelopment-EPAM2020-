package ht1

import (
	"sort"
)

// Median : This method is used to calculate median statistic of array
func Median(i []int) float64 {
	sort.Ints(i)
	if len(i)%2 != 0 {
		return float64(i[(len(i)-1)/2])
	}
	return float64((i[(len(i)-1)/2] + i[(len(i)-1)/2+1])) / 2
}

// Call : This method is used for check ht1 in 'main.go'
func Call(numbers []int) float64 {
	return Median(numbers)
}
