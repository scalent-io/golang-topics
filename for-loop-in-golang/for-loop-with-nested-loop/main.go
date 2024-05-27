package main

import "fmt"

func main() {
	size := 5

	// Outer loop for the rows
	for row := 1; row <= size; row++ {
		// Inner loop for the columns
		for column := 1; column <= size; column++ {
			fmt.Printf("%4d", row*column) // Print the product of i and j
		}
		fmt.Println() // Move to the next line after each row
	}
}
