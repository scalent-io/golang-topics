package main

import "fmt"

func main() {
	shirtColor := "WHITE"

	if shirtColor == "WHITE" {
		fmt.Println("The color of the shirt is white")
	} else if shirtColor == "BLACK" {
		fmt.Println("The color of the shirt is black")
	} else if shirtColor == "RED" {
		fmt.Println("The color of the shirt is red")
	} else if shirtColor == "BROWN" {
		fmt.Println("The color of the shirt is brown")
	}
}
