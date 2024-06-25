package main

import "fmt"

var memo = map[int]int{}

func fibonacciMemoized(n int) int {
	if v, ok := memo[n]; ok {
		return v
	}
	if n <= 1 {
		return n
	}
	result := fibonacciMemoized(n-1) + fibonacciMemoized(n-2)
	memo[n] = result
	return result
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(fibonacciMemoized(i))
	}
}
