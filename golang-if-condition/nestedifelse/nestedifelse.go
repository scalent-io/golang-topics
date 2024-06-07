package main

import "fmt"

func main() {

	//find employee

	a := 19 //number1
	b := 21 // number2
	c := 28 // number3

	if a > b {
		if a > c {
			fmt.Println("Largest number is ", a)
		} else {
			fmt.Println("Largest number is ", c)
		}
	} else if b > a {
		if b > c {
			fmt.Println("Largest number is ", b)
		} else {
			fmt.Println("Largest number is ", c)
		}
	} else if c > a {
		if c > b {
			fmt.Println("Lrgest number is ", c)
		} else {
			fmt.Println("Largest number is", b)
		}
	}
}
