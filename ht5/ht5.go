package ht5

import "sort"

func printSorted(dict map[int]string) []string {
	var resultSlice []string
	var keys []int
	for key := range dict {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	for _, key := range keys {
		resultSlice = append(resultSlice, dict[key])
	}
	return resultSlice
}

// Call : This method is used for check ht5 in 'main.go'
func Call(dict map[int]string) []string {
	return printSorted(dict)
}
