package main

import (
	"fmt"
)

type EmployeeContact struct {
	Name  string
	Phone string
}

func main() {

	// Create an instance of EmployeeContact
	employee1 := EmployeeContact{"Scalent_Employee_1", "1234567890"}

	// Print the original struct instance
	fmt.Println("Original employee1:", employee1) // Outputs: {Scalent_Employee_1 1234567890}

	// Create a pointer to the struct instance
	employee1Pointer := &employee1

	// Modify the struct fields using the pointer
	employee1Pointer.Phone = "0987654321"

	// Print the modified struct instance
	fmt.Println("Modified employee1:", employee1) // Outputs: {Scalent_Employee_1 0987654321}
}
