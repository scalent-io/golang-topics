package main

import "fmt"

func multiplyMatrices(a, b [2][2]int) [2][2]int {
	result := [2][2]int{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				result[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return result
}

func matrixExponentiation(matrix [2][2]int, n int) [2][2]int {
	result := [2][2]int{{1, 0}, {0, 1}}
	for n > 0 {
		if n&1 == 1 {
			result = multiplyMatrices(result, matrix)
		}
		matrix = multiplyMatrices(matrix, matrix)
		n >>= 1
	}
	return result
}

func fibonacciMatrixExponentiation(n int) int {
	if n <= 1 {
		return n
	}
	matrix := [2][2]int{{1, 1}, {1, 0}}
	result := matrixExponentiation(matrix, n-1)
	return result[0][0]
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciMatrixExponentiation(i))
	}
}
