package main

import "fmt"

func main() {

	name := "Alice"
	age := 30
	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(formattedString)
	// Output: Name: Alice, Age: 30

}
