package main

import "fmt"

func main() {
	// Loop to find the first number greater than 5
	for num := 1; num <= 10; num++ {
		if num > 5 {
			fmt.Printf("Breaking the loop at number: %d\n", num)
			break // Exit the loop
		}
	}
}
