package main

import (
	"fmt"

	"./ht1"

	"./ht2"
)

func main() {
	fmt.Println("Home task 1: ", ht1.Call([]int{0, 3, 2, 7, 16, 5, 1, 4}))

	endPoint, perimeter, area := ht2.Call()
	fmt.Println("Home task 2:\n\tendPoint=", endPoint, "\n\tperimeter=", perimeter, "\n\tarea=", area)
}
