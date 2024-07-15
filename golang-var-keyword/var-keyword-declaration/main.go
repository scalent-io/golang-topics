package main

import "fmt"

func main() {
	// Variables with explicit type declarations
	var age int = 30
	var name string = "James"
	var isLoggedIn bool = false

	// Variables with type inference
	var x = 10                    // Inferred type: int
	var message = "Hello, world!" // Inferred type: string
	var success = true            // Inferred type: bool

	// Multiple variables of the same type in a single line
	var width, height int = 100, 200

	// Print the values of the variables
	fmt.Println("Explicit type declarations:")
	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Is Logged In:", isLoggedIn)

	fmt.Println("\nType inference:")
	fmt.Println("x:", x)
	fmt.Println("Message:", message)
	fmt.Println("Success:", success)

	fmt.Println("\nMultiple variables:")
	fmt.Println("Width:", width, "Height:", height)
}
