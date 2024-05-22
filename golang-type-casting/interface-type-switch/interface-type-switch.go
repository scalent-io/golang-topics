package main

import (
	"fmt"
)

func main() {
	var i interface{}

	// Assign different types to the empty interface
	i = 42
	checkType(i)

	i = "hello"
	checkType(i)

	i = 3.14
	checkType(i)

	i = true
	checkType(i)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int: %d\n", v)
	case string:
		fmt.Printf("i is a string: %s\n", v)
	case float64:
		fmt.Printf("i is a float64: %f\n", v)
	case bool:
		fmt.Printf("i is a bool: %t\n", v)
	default:
		fmt.Printf("i is of an unknown type: %T\n", v)
	}
}
