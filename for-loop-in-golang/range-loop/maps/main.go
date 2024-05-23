package main

import "fmt"

func main() {
	user := map[string]string{
		"Name":      "Rob Pike",
		"Address":   "Googleplex.",
		"IsCreator": "Yes",
	}
	fmt.Println("User details are as follows")
	for key, value := range user {
		fmt.Println(key, ":", value)
	}
}
