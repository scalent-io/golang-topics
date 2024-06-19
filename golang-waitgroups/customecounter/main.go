package main

import (
	"fmt"
	"sync/atomic"
)

type custWaitGroup struct {
	count int64
}

func (wg *custWaitGroup) Add(n int64) {
	atomic.AddInt64(&wg.count, n)
}

func (wg *custWaitGroup) Done() {
	wg.Add(-1)
	if atomic.LoadInt64(&wg.count) < 0 {
		panic("negative wait group counter")
	}
}

func (wg *custWaitGroup) Wait() {
	for atomic.LoadInt64(&wg.count) != 0 {
		continue
	}
}

func main() {
	var wg custWaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("GoroutinesTwo")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("GoroutineOne")
	}()
	wg.Wait()
	fmt.Println("All goroutines are completed their execution")
}
