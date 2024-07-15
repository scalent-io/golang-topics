package main

import "fmt"

// Define a struct Circle
type Circle struct {
	radius float64
}

func main() {
	// Create a variable of Circle struct type
	var c Circle = Circle{radius: 5}

	// Access and print the radius
	fmt.Println("Radius of circle:", c.radius)
}
