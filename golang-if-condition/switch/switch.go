package main

import (
	"fmt"
)

func main() {
	var a, b float64
	var operation string

	// Get user input
	fmt.Print("Enter the first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter the second number: ")
	fmt.Scanln(&b)
	fmt.Print("Enter the operation (+, -, *, /): ")
	fmt.Scanln(&operation)

	switch operation {
	case "+":
		fmt.Printf("The result of %.2f + %.2f is: %.2f\n", a, b, a+b)
	case "-":
		fmt.Printf("The result of %.2f - %.2f is: %.2f\n", a, b, a-b)
	case "*":
		fmt.Printf("The result of %.2f * %.2f is: %.2f\n", a, b, a*b)
	case "/":
		if b != 0 {
			fmt.Printf("The result of %.2f / %.2f is: %.2f\n", a, b, a/b)
		} else {
			fmt.Println("Error: Division by zero is not allowed.")
		}
	default:
		fmt.Println("Error: Invalid operation. Please enter one of +, -, *, /.")
	}
}
