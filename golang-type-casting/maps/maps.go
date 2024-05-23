package main

import (
	"fmt"
)

func main() {
	intMap := map[string]int{"Number1": 1000, "Number2": 2000, "Number3": 3000}
	floatMap := make(map[string]float64)

	for k, v := range intMap {
		floatMap[k] = float64(v)
	}

	fmt.Println("Converted float map:", floatMap)
}
