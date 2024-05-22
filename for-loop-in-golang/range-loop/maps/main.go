package main

import "fmt"

func main() {
	user := map[string]string{
		"name":        "Sachin Tendulkar",
		"phoneNumber": "7111111112",
		"Address":     "At post yawatmal.",
	}
	for key, value := range user {
		fmt.Println(key, value)
	}
}
