package main

import (
	"fmt"
	"math/rand"
)

func main() {
	winningNumber := 7
	for {
		number := rand.Intn(10)
		if number == winningNumber {
			fmt.Println("You won the lottery! Exiting the loop...")
			break
		}
	}
	fmt.Println("Loop exited.  Enjoy your winnings!")
}
