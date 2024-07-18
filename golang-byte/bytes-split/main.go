package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("apple,banana,orange")
	parts := bytes.Split(data, []byte(","))
	fmt.Println(parts) // Output: [[97 112 112 108 101] [98 97 110 97 110 97] [111 114 97 110 103 101]]
}
