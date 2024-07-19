package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("example")
	fmt.Println(bytes.Count(data, []byte("e"))) // Output: 2
}
