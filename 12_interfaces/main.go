package main

import (
	"fmt"
	"math"
)

// Define interface type

type Shape interface {
	area() float64
}

// Structs of shapes
type Circle struct {
	radius float64
}

type Rectangle struct {
	length, breadth float64
}

func (r Rectangle) area() float64 {
	return r.length * r.breadth
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	c := Circle{radius: 5}
	fmt.Println(getArea(c))
	r := Rectangle{length: 10, breadth: 20}
	fmt.Println(getArea(r))
}
