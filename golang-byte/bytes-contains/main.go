package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Define the byte slice and the sub-slice to find
	text := []byte("Example")  // Using string literal to create byte slice
	subText := []byte("ample") // Using string literal to create byte slice for the sub-slice

	// Check if the byte slice contains the sub-slice
	contains := bytes.Contains(text, subText)
	if contains {
		fmt.Println("The byte slice contains the sub-slice.")
	} else {
		fmt.Println("The byte slice does not contain the sub-slice.")
	}
}
