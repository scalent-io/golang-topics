package main

import "fmt"

func main() {
	str := "🙏🙏"

	// Iterate over the string using range
	for index, runeValue := range str {
		fmt.Printf("Character at index %d is %c\n", index, runeValue)
	}
}
