package main

import "fmt"

func main() {
	//-------send values to ch1 in loop and read values in goroutine 1 in loop
	ch1 := make(chan int)
	go goRoutine(ch1)

	for counter := 1; counter <= 5; counter++ {
		ch1 <- counter
	}
}

func goRoutine(ch chan int) {
	for channelValue := range ch {
		fmt.Println("Current value in channel is: ", channelValue)
	}
}
