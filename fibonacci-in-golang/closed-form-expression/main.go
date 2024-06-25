package main

import (
	"fmt"
	"math"
)

func fibonacciBinet(n int) int {
	phi := (1 + math.Sqrt(5)) / 2
	psi := (1 - math.Sqrt(5)) / 2
	return int((math.Pow(phi, float64(n)) - math.Pow(psi, float64(n))) / math.Sqrt(5))
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fibonacciBinet(i))
	}
}
