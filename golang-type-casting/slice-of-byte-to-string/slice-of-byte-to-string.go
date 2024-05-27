package main

import "fmt"

func main() {
	byteSlice := []byte{71, 111, 112, 104, 101, 114, 115, 33} // Represents "Gophers!"
	str := string(byteSlice)
	fmt.Println(str) // Output: Gophers!
}
