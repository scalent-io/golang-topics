//problem statement : convert float64 to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var percentageValue float64 = 12.5
	var percentageValueString string = strconv.FormatFloat(percentageValue, 'f', 2, 64)
	message := percentageValueString + " percentage is equal to 1/8 fraction"
	fmt.Println(message)
}

//output : 12.50 percentage is equal to 1/8 fraction
