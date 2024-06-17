package main

import (
	"fmt"
	"time"
)

func main() {
	// a function call with no go keyword
	Greeting("John")
	Greeting("James")
	Greeting("Michael")

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
