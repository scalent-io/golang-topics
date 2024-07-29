package main

import "fmt"

func main() {

	var x int = 42  // Create an integer variable x with a value of 42
	var p *int = &x // Create a pointer variable p that stores the address of x

	fmt.Println("Value of x:", x)    // Output: 42
	fmt.Println("Address of x:", &x) // Output: address of x: 0xc0000120b0
	fmt.Println("Value of p:", p)    // Output: value of p: 0xc0000120b0 (same as the address of x)
	fmt.Println("Value at p:", *p)   // Output: 42 (value stored at the memory address p points to)
}
