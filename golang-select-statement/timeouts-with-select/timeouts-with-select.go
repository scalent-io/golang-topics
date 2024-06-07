package main

import (
	"fmt"
	"time"
)

func main() {

	taskChannel := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		taskChannel <- "Task completed"
	}()

	select {
	case msg := <-taskChannel:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout: Task took too long")
	}
}
