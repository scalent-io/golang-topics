//Problem statement: convert a integer to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create the weight integer
	var weight int = 60

	//Create the weight string
	var weightString string = strconv.Itoa(weight)

	// Create the message
	message := "My weight is " + weightString + " kg"

	// Print the message
	fmt.Printf("%v\nDatatype of message is %T\n", message, message)
}

/* output:
My weight is 60 kg
Datatype of message is string*/
