package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
			return
		case msg2 := <-ch2:
			fmt.Println(msg2)
			return
		default:
			fmt.Println("No messages received yet, doing other work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
