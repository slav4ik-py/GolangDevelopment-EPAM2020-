package ht4

func max(array []string) string {
	var maxLength int
	var stringWithMaxLength string
	for _, i := range array {
		if len(i) > maxLength {
			maxLength = len(i)
			stringWithMaxLength = i
		}
	}
	return stringWithMaxLength
}

func reverse(array []int64) []int64 {
	var resultSlice []int64
	for i := len(array) - 1; i >= 0; i-- {
		resultSlice = append(resultSlice, array[i])
	}
	return resultSlice
}

// CallMax : This method is used for check 'max' of ht3 in 'main.go'
func CallMax(array []string) string {
	return max(array)
}

// CallReverse : This method is used for check 'reverse' of ht3 in 'main.go'
func CallReverse(array []int64) []int64 {
	return reverse(array)
}
