package main

import (
	"fmt"
)

func main() {
	var intArray = [3]int{1, 2, 3}
	var floatArray [3]float64

	for i, v := range intArray {
		floatArray[i] = float64(v)
	}

	fmt.Println("Converted float array:", floatArray)
}
