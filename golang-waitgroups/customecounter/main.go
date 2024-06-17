package main

import (
	"fmt"
	"sync/atomic"
)

type myWaitGroup struct {
	count int64
}

func (wg *myWaitGroup) Add(n int64) {
	atomic.AddInt64(&wg.count, n)
}

func (wg *myWaitGroup) Done() {
	wg.Add(-1)
	if atomic.LoadInt64(&wg.count) < 0 {
		panic("negative wait group counter")
	}
}

func (wg *myWaitGroup) Wait() {
	for atomic.LoadInt64(&wg.count) != 0 {
		continue
	}
}

func main() {
	var wg myWaitGroup
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
