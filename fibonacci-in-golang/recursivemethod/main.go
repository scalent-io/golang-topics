package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 10
	fmt.Println("Fibonacci sequence:")
	for i := 0; i < n; i++ {
		fmt.Printf("F(%d) = %d\n", i, fibonacci(i))
	}
}
