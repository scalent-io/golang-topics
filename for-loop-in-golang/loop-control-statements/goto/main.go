package main

import "fmt"

func main() {
	num := 0
	for {
		if num == 5 {
			goto endloop // Jump to the end of the loop
		}
		fmt.Println(num)
		num++
	}
endloop:
	fmt.Println("Loop ended")
}
