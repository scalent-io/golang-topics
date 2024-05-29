package main

import (
	"fmt"
)

func main() {
	intSlice := []int{1000, 2000, 3000}
	floatSlice := make([]float64, len(intSlice))

	for i, v := range intSlice {
		floatSlice[i] = float64(v)
	}

	fmt.Println("Converted float slice:", floatSlice)
}
