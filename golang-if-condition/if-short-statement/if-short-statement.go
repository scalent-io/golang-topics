package main

import (
	"fmt"
	"os"
)

func main() {
	//you should mention the file path to open the file
	if file, err := os.Open("/home/scalent/Downloads/EnglishPTP.pdf"); err == nil {
		fmt.Println("File opened successfully")
		// Make sure to close the file
		file.Close()
	} else {
		fmt.Println("Error opening file:", err)
	}
}
