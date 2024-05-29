package main

import (
	"fmt"
)

func main() {
	var greetings interface{} = "Greetings of the day!"

	// Type assertion to convert interface{} to string
	s, ok := greetings.(string)
	if ok {
		fmt.Println("The string is:", s)
	} else {
		fmt.Println("Type assertion failed")
	}
}
