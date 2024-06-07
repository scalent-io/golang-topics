package main

import "fmt"

func main() {
	byteSlice := []byte{72, 101, 108, 108, 111} // Represents "Hello"
	str := string(byteSlice)
	fmt.Println(str) // Output: Hello
}
