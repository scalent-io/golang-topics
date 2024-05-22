//problem statement : convert float64 to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create the weight float
	var weight float64 = 52.50

	// Create the weight string
	var weightString string = strconv.FormatFloat(weight, 'f', 2, 64)

	// Create the message
	message := "My weight is " + weightString + " kg"

	// Print the message
	fmt.Printf("%v\nDatatype of message is %T\n", message, message)
}

/*
output:
My weight is 50.5 kg
Datatype of message is string
*/
