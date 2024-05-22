package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Create the circleOutsideDiameter string
	var circleOutsideDiameterString string = "35.30"

	// Convert the circleOutsideDiameterString to a float64 first
	circleOutsideDiameter64, err := strconv.ParseFloat(circleOutsideDiameterString, 32)
	if err != nil {
		fmt.Println("Error converting string to float64:", err)
		return
	}

	if circleOutsideDiameter64 > math.MaxFloat32 || circleOutsideDiameter64 < -math.MaxFloat32 {
		fmt.Println("Value is out of range for float32")
		return
	}

	// Cast the float64 to float32
	circleOutsideDiameter := float32(circleOutsideDiameter64)

	// Print the converted float32
	fmt.Printf("%v\nDatatype of circleOutsideDiameter is %T\n", circleOutsideDiameter, circleOutsideDiameter)
}

/*
output :
30.3
Datatype of circleOutsideDiameter is float32
*/
