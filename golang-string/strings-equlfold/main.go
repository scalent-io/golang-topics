package main

import (
	"fmt"
	"strings"
)

func main() {

	equal := strings.EqualFold("GoLang", "golang")
	if equal {
		fmt.Println("The strings are equal, ignoring case")
	} else {
		fmt.Println("The strings are not equal")
	}

}
