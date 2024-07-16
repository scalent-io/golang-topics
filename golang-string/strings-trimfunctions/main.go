package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("!!!hello, world!!!", "!"))          // Output: hello, world
	fmt.Println(strings.TrimSpace(" \t\n hello, world \n\t\r\n")) // Output: hello, world
	fmt.Println(strings.TrimPrefix("hello, world", "hello"))      // Output: , world
	fmt.Println(strings.TrimSuffix("hello, world", "world"))      // Output: hello,
	fmt.Println(strings.TrimLeft("@@@@cool, world!!!", "@"))      // Output: cool, world!!!
	fmt.Println(strings.TrimRight("!!!hello, world!!!", "!"))     // Output: !!!hello, world
}
