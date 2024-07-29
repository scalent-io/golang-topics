package main

import "fmt"

// Define a struct type
type Rectangle struct {
	Width, Height float64
}

// Function to resize a rectangle by modifying its fields using a pointer
func resizeRectangle(r *Rectangle, width, height float64) {
	r.Width = width
	r.Height = height
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	fmt.Println("Before resizing:", rect) // Outputs: Before resizing: {5 10}

	// Pass the address of rect to the function
	resizeRectangle(&rect, 15, 20)

	fmt.Println("After resizing:", rect) // Outputs: After resizing: {15 20}
}
