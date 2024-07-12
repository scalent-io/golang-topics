package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello,world,how,are,you"
	sep := ","
	result := strings.SplitAfter(str, sep)
	fmt.Println(result)
}

//[hello, world, how, are, you]
