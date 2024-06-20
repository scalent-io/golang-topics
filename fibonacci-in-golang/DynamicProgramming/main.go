package main

import "fmt"

func fibonacciDynamicProgramming(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fibonacciDynamicProgramming(i))
	}
}
