package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("Goroutines started")

	wg.Add(3) // we have 3 goroutines , so we added 3 to add()

	go Greeting("John", &wg)
	go Greeting("James", &wg)
	go Greeting("Michael", &wg)

	wg.Wait()

	fmt.Println("All goroutines completed")
}

func Greeting(name string, wg *sync.WaitGroup) {
	defer wg.Done()
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
