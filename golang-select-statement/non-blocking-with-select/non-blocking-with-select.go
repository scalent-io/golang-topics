package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	select {
	case value := <-ch:
		fmt.Println("Received value:", value)
	default:
		fmt.Println("No value received")
	}
}
