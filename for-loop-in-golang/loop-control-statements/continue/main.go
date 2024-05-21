package main

import "fmt"

func main() {
	// Loop to print all numbers between 1 and 10 except 5
	for num := 1; num <= 5; num++ {
		if num == 3 {
			continue // Skip 5th iteration
		}
		fmt.Println(" Iteration number : ", num)
	}
}
