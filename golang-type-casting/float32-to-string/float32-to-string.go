//problem statement : convert the float32 to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create the circleOutsideDiameter float32
	var circleOutsideDiameter float32 = 30.35

	// Convert the float32 to a string
	circleOutsideDiameterString := strconv.FormatFloat(float64(circleOutsideDiameter), 'f', 2, 32)

	// Print the converted string
	fmt.Printf("%v\nDatatype of circleOutsideDiameterString is %T\n", circleOutsideDiameterString, circleOutsideDiameterString)
}

/*
output :
30.30
Datatype of circleOutsideDiameterString is string
*/
