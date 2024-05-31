package main

import "fmt"

func main() {
	ageOfPassenger := 55

	if ageOfPassenger < 0 || ageOfPassenger > 125 {
		fmt.Println("Invalid age")
	} else if ageOfPassenger > 0 && ageOfPassenger <= 5 {
		fmt.Println("There is no ticket for children below 5 years")
	} else if ageOfPassenger > 5 && ageOfPassenger <= 60 {
		fmt.Println("There is a ticket for children above 5 years and people below 60 years")
	} else {
		fmt.Println("There is no ticket for senior citizens")
	}
}
