//problem statement : convert string to float64

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Create the circleOutsideDiameter string
	var circleOutsideDiameterString string = "30.35"

	// Convert the circleOutsideDiameterString to a float64
	circleOutsideDiameter, err := strconv.ParseFloat(circleOutsideDiameterString, 64)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	// Print the converted float64
	fmt.Printf("%v\nDatatype of circleOutsideDiameter is %T\n", circleOutsideDiameter, circleOutsideDiameter)
}

/*
output :
30.3
Datatype of circleOutsideDiameter is float64
*/
