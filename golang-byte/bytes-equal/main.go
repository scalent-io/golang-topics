package main

import (
	"bytes"
	"fmt"
)

func main() {
	byteSlice1 := []byte("Golang")
	byteSlice2 := []byte("Golang")
	fmt.Println(bytes.Equal(byteSlice1, byteSlice2)) // Output: true
}
