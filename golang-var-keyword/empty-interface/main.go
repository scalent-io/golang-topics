package main

import "fmt"

func main() {
	// Declare a variable of empty interface type
	var x interface{}

	// Assign different values to the empty interface variable
	x = 42 // int
	fmt.Println("Type:", getType(x), ", Value:", x)

	x = "hello" // string
	fmt.Println("Type:", getType(x), ", Value:", x)

	x = true // bool
	fmt.Println("Type:", getType(x), ", Value:", x)

	x = struct{}{} // empty struct
	fmt.Println("Type:", getType(x), ", Value:", x)
}

// Function to get the type of the interface value
func getType(i interface{}) string {
	return fmt.Sprintf("%T", i)
}
