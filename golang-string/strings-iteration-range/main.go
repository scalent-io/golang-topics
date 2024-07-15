package main

import (
	"fmt"
)

func main() {
	str := "Hello, Golang"
	// Iterating over characters using range
	for index, runeValue := range str {
		fmt.Printf("Character at index %d is %c\n", index, runeValue)
	}
}
