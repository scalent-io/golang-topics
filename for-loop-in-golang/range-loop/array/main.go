package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red", "yellow"}

	for index, value := range colors {
		fmt.Printf("index: %d, value: %v\n", index, value)
	}
}
