package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "🙏🙏"

	// Counting runes in the string
	runeCount := utf8.RuneCountInString(str)
	for i := 0; i < runeCount; i++ {
		// Decode the next rune in the string
		runeValue, size := utf8.DecodeRuneInString(str)

		// Print the index and the decoded rune
		fmt.Printf("index: %d, character: %c\n", i, runeValue)

		// Move to the next rune in the string by slicing out the bytes of the decoded rune
		str = str[size:]
	}
}
