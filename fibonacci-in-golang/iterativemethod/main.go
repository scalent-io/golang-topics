package main

import "fmt"

func fibonacciCalculator(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fibonacciCalculator(i))
	}
}
