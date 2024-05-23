package main

import "fmt"

func main() {
	signals := []string{"green", "yellow", "red"} // Example signals

	// Loop to check each signal
	for i, signal := range signals {
		if signal == "red" {
			fmt.Printf("Breaking the loop at signal: %s (index: %d)\n", signal, i)
			break // Exit the loop if signal is red
		}
	}
}
