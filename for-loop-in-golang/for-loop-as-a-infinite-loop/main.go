package main

import (
	"fmt"
)

func main() {
	counter := 0
	for {
		fmt.Println("Hello, World!")
		counter++
		if counter >= 5 {
			break // Exit the loop after 5 iterations
		}
	}
	fmt.Println("Loop exited.")
}
