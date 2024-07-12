package main

import (
	"fmt"
	"strings"
)

func main() {
	isPrefix := strings.HasPrefix("hello, world", "hello")
	if !isPrefix {
		fmt.Println("The string does not contain the substring at start")
	} else {
		fmt.Println(isPrefix)
	}

	isSuffix := strings.HasSuffix("hello, world", "world")
	if !isSuffix {
		fmt.Println("The string does not contain the substring at end")
	} else {
		fmt.Println(isSuffix)
	}

}

// Output: true
// Output: true
