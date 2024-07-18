package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("example")
	fmt.Println(bytes.Index(data, []byte("mp"))) // Output: 3
}
