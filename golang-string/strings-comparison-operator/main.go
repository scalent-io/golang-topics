package main

import "fmt"

func main() {
	str1 := "apple"
	str2 := "banana"

	// Direct comparison
	if str1 == str2 {
		fmt.Println("Strings are equal")
	} else if str1 < str2 {
		fmt.Println("str1 is less than str2")
	} else {
		fmt.Println("str1 is greater than str2")
	}
}
