package main

import "fmt"

// Define a struct type
type Rectangle struct {
	Width, Height float64
}

// Method with a value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method with a pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func main() {
	// Create an instance of Rectangle
	rect := Rectangle{Width: 5, Height: 10}

	// Print the original area
	fmt.Println("Original Area:", rect.Area()) // Outputs: Original Area: 50

	// Scale the rectangle
	rect.Scale(2)

	// Print the new area after scaling
	fmt.Println("New Area:", rect.Area()) // Outputs: New Area: 200
}
