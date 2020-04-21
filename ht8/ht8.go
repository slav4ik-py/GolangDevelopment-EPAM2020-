package main

import "fmt"

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	lengthOfSide float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.lengthOfSide * s.lengthOfSide
}

func (s Square) perimeter() float64 {
	return s.lengthOfSide * 4
}

func (c Circle) area() float64 {
	return 3.141592 * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * 3.141592 * c.radius
}

func main() {
	var s Figure = Square{5}
	var c Figure = Circle{3}
	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())
}
