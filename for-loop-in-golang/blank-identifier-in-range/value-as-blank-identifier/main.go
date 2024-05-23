package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red", "white"}

	// Iterate over each element in the colors slice, ignoring the index
	for index, _ := range colors {
		fmt.Println("Color index is :", index)
	}
}
