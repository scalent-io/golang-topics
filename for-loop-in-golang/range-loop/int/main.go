package main

import "fmt"

func main() {
	for i := range 10 {
		fmt.Println(10 - i)
	}
	fmt.Println("go1.22 has lift-off!")
}
