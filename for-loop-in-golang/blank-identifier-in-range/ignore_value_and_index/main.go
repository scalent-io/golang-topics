package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red", "white"}

	// Iterate over each element in the colors slice, ignoring both index and value
	for range colors {
		fmt.Println("Same color!")
	}
}
