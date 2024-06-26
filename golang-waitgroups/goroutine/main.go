package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Goroutines started")
	// a function call with go keyword is function called in newly invoked goroutine
	go Greeting("John")
	go Greeting("James")
	go Greeting("Michael")

	fmt.Println("All goroutines completed")
}

func Greeting(name string) {
	now := time.Now()
	hour := now.Hour()
	switch {
	case hour >= 12 && hour < 17: // 12:00 PM to 5:00 PM
		fmt.Printf("Good afternoon, %v!\n", name)
	case hour < 12: // before 12:00 PM
		fmt.Printf("Good morning, %v!\n", name)
	case hour >= 17 && hour < 22: // after 5:00 PM
		fmt.Printf("Good evening, %v!\n", name)
	default:
		fmt.Printf("Good night, %v!\nIt's time to sleep!\n", name)
	}
}
