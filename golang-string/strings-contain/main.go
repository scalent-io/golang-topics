package main

import (
	"fmt"
	"strings"
)

func main() {
	exist := strings.Contains("hello, world", "world")
	if !exist {
		fmt.Println("String does not contain the substring")
	}
	fmt.Println(exist)
}

// output : true
