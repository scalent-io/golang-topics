package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Replace("hello, Gophers!", "o", "0", -1))
	// Output: hell0, G0phers!

}
