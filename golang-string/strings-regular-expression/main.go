package main

import (
	"fmt"
	"regexp"
)

func main() {

	// example 1 checking pattern existance
	// Define a regular expression pattern to match "hello"
	pattern := regexp.MustCompile(`hello`)

	// Example strings to test
	testString1 := "say hello to everyone"
	testString2 := "goodbye for now"

	// Check if each string contains "hello" and print results
	fmt.Printf("'%s' contains 'hello': %v\n", testString1, pattern.MatchString(testString1))
	fmt.Printf("'%s' contains 'hello': %v\n", testString2, pattern.MatchString(testString2))

	// Example 2 : text containing programming symbols
	text := "Let's extract symbols like +, -, *, /, =, and ! from this text."

	// Define a regular expression pattern to match programming symbols
	symbolPattern := regexp.MustCompile(`[+\-*/=!]+`)

	// Find all symbols in the text using the pattern
	symbols := symbolPattern.FindAllString(text, -1)

	// Print the extracted symbols
	fmt.Println("Extracted symbols:")
	for _, symbol := range symbols {
		fmt.Println(symbol)
	}

}
