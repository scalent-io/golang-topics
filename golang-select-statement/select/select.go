package main

import (
	"fmt"
	"time"
)

func evenNumberSum(from, to int, ch chan int) {
	totalEvenSum := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			totalEvenSum += i
		}
	}
	ch <- totalEvenSum
}
func squareNumberSum(from, to int, ch chan int) {
	totalSquareSum := 0
	for i := from; i <= to; i++ {
		if i%2 == 0 {
			totalSquareSum += i * i
		}
	}
	ch <- totalSquareSum
}

func main() {
	evenNumberChannel := make(chan int)
	squareNumberChannel := make(chan int)

	go evenNumberSum(0, 100, evenNumberChannel)
	go squareNumberSum(0, 100, squareNumberChannel)

	timeout := time.After(1 * time.Second) // Set a timeout for 1 seconds

	for {
		select {
		case totalEvenSum := <-evenNumberChannel:
			fmt.Println("Total Even Sum:", totalEvenSum)
			return
		case totalSquareSum := <-squareNumberChannel:
			fmt.Println("Total Square Sum:", totalSquareSum)
			return
		case <-timeout:
			fmt.Println("Timeout! No data received.")
			return
		default:
			fmt.Println("Presently no data is received ")
			time.Sleep(5 * time.Millisecond)
		}
	}
}
