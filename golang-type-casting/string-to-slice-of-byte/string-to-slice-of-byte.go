package main

import (
	"fmt"
)

func main() {
	str := "Gophers!"
	byteSlice := []byte(str)
	fmt.Println(byteSlice)
}

//output : [71 111 112 104 101 114 115 33]
