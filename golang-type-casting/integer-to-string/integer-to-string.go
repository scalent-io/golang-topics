//Problem statement: convert a integer to string

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var secondsInOneMinute int = 60
	var secondsInOneMinuteString string = strconv.Itoa(secondsInOneMinute)
	message := "One minute consists of " + secondsInOneMinuteString + " seconds"
	fmt.Println(message)
} // output : One minute consists of 60 seconds
