package main

import "fmt"

func main() {
	var a int = 10  // Declare an integer variable a with value 10
	var p *int = &a // Declare a pointer variable p and assign it the address of a

	// Dereference p to access the value at the address it points to
	fmt.Println("Value at pointer p:", *p) // Outputs: 10 (value stored at the address p points to)

	// Modify the value at the address p points to
	*p = 20
	fmt.Println("New value of a:", a) // Outputs: 20 (a has been updated through pointer p)
}
