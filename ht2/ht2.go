package ht2

// Point : This structure is used to define Point type
type Point struct {
	x, y int
}

// Square : This structure is used to define Square type
type Square struct {
	start Point
	a     uint
}

// End : This method is used to find end point of square
func (s Square) End() Point {
	return Point{s.start.x + int(s.a), s.start.y - int(s.a)}
}

// Perimeter : This method is used to find a perimeter of square
func (s Square) Perimeter() uint {
	return s.a * 4
}

// Area : This method is used to find an area of square
func (s Square) Area() uint {
	return s.a * s.a
}

// Call : This method is used for check ht2 in 'main.go'
func Call() (Point, uint, uint) {
	s := Square{Point{1, 1}, 5}
	return s.End(), s.Perimeter(), s.Area()
}
