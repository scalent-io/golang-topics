package main

import (
	"fmt"
)

func main() {
	intMap := map[string]int{"a": 1, "b": 2, "c": 3}
	floatMap := make(map[string]float64)

	for k, v := range intMap {
		floatMap[k] = float64(v)
	}

	fmt.Println("Converted float map:", floatMap)
}
