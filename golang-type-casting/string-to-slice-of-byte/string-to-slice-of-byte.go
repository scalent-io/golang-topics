package main

import (
	"fmt"
)

func main() {
	str := "Hello, world!"
	byteSlice := []byte(str)
	fmt.Println(byteSlice)
}

//output : [72 101 108 108 111 44 32 119 111 114 108 100 33]
