package ht3

func average(array [6]int) float64 {
	var sum int
	for _, i := range array {
		sum += i
	}
	return float64(sum) / float64(len(array))
}

// Call : This method is used for check ht3 in 'main.go'
func Call(array [6]int) float64 {
	return average(array)
}
