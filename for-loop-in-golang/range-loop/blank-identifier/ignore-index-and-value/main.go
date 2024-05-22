package main

import "fmt"

func main() {
	colors := []string{"blue", "green", "red", "yellow"}

	for range colors {
		fmt.Println("Just looping over")
	}
}
