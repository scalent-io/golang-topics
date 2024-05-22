package main

import "fmt"

func main() {
	fruits := []string{"apple", "banana", "cherry"}

	for index, _ := range fruits {
		fmt.Println(index)
	}
}
