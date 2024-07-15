package main

import (
	"fmt"
	"strings"
)

func main() {
	result := strings.Compare("Gopher", "Gopher")
	switch result {
	case 0:
		fmt.Println("Strings are equal")
	case -1:
		fmt.Println("First string is less than the second")
	case 1:
		fmt.Println("First string is greater than the second")
	}
}
