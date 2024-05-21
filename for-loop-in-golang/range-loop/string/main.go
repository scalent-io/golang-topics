package main

import "fmt"

func main() {
	greeting := "Good Morning!"

	for index, runeValue := range greeting {
		fmt.Printf("Index: %d, Character: %c\n", index, runeValue)
	}
}
