package main

import (
	"fmt"
	"math"
)

// Define an interface
type Shape interface {
	area() float64
}

// Define a struct Circle
type Circle struct {
	radius float64
}

// Implementing Shape interface for Circle
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	// Create an instance of Circle
	circle := Circle{radius: 5}

	// Create a variable of Shape interface type & assign the Circle instance to the Shape interface variable
	var shape Shape = circle

	// Call area method on Circle using Shape interface
	area := shape.area()

	// Print area of circle
	fmt.Printf("Area of circle: %.2f\n", area)
}
