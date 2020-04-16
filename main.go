package main

import (
	"fmt"

	"./ht1"

	"./ht2"

	"./ht3"

	"./ht4"

	"./ht5"
)

func main() {
	fmt.Println("Home task 1: ", ht1.Call([]int{0, 3, 2, 7, 16, 5, 1, 4}))

	endPoint, perimeter, area := ht2.Call()
	fmt.Println("Home task 2:\n\tendPoint=", endPoint, "\n\tperimeter=", perimeter, "\n\tarea=", area)

	fmt.Println("Home task 3: ", ht3.Call([6]int{0, 3, 2, 7, 16, 5}))

	fmt.Println("Home task 4:\n\tmax:", ht4.CallMax([]string{"one", "two"}), "\n\treverse:", ht4.CallReverse([]int64{1, 2, 5, 15}))

	fmt.Println("Home task 5: ", ht5.Call(map[int]string{10: "aa", 0: "bb", 500: "cc"}))
}
