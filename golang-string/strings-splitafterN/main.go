package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "apple-orange-pear-grape-melon"
	// Split after each '-' with a maximum of 3 splits
	parts := strings.SplitAfterN(str, "-", 3)
	fmt.Println("Split parts:", parts)
}

//output : Split parts: [apple- orange- pear- grape-melon]
