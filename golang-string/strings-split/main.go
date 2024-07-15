package main

import (
	"fmt"
	"strings"
)

func main() {
	sliceOfSubstrings := strings.Split("One,Two,Three", ",")
	fmt.Println(sliceOfSubstrings) // Output: [One Two Three]

	string := strings.Join([]string{"One", "Two", "Three"}, "-")
	fmt.Println(string) // One-Two-Three
}
