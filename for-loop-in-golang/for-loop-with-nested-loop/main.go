package main

import "fmt"

func main() {
	size := 10

	// Outer loop for the rows
	for i := 1; i <= size; i++ {
		// Inner loop for the columns
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j) // Print the product of i and j
		}
		fmt.Println() // Move to the next line after each row
	}
}
