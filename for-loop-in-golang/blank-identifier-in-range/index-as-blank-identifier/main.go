package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red", "white"}

	// Iterate over each element in the colors slice, ignoring the index
	for _, value := range colors {
		fmt.Println("Color is :", value)
	}
}
