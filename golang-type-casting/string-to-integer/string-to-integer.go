// Problem statement: convert a string to int

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create the mobileNumber string
	var mobileNumberString string = "1234567890.25"

	// Convert the mobileNumberString to an int
	mobileNumber, err := strconv.Atoi(mobileNumberString)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}

	// Print the converted integer
	fmt.Printf("%v\nDatatype of mobileNumber is %T\n", mobileNumber, mobileNumber)
}

/*
output :
1234567890
Datatype of mobileNumber is int
*/
