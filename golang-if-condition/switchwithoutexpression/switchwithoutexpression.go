package main

import "fmt"

func main() {
	age := 30

	switch {
	case age < 0:
		fmt.Println("Invalid age")
	case age <= 5:
		fmt.Println("There is no ticket for children below 5 years")
	case age <= 60:
		fmt.Println("There is a ticket for children above 5 years and people below 60 years")
	default:
		fmt.Println("There is no ticket for senior citizens")
	}
}
