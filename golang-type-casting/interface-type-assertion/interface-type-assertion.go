package main

import (
	"fmt"
)

func main() {
	var i interface{} = "hello"

	// Type assertion to convert interface{} to string
	s, ok := i.(string)
	if ok {
		fmt.Println("The string is:", s)
	} else {
		fmt.Println("Type assertion failed")
	}
}
