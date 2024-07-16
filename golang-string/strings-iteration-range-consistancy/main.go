package main

import "fmt"

func main() {

	str := "🙏🙏"
	runestr := []rune(str)
	// Iterating over characters and printing Unicode code points
	for index, runeValue := range runestr {
		fmt.Printf("value:%c\tindexed at %d\n", runeValue, index)
	}
}
