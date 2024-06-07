package main

import (
	"fmt"
)

func main() {
	signals := []string{"green", "yellow", "red"} // Example signals

	// Loop to check each signal
	for i, signal := range signals {
		if signal == "green" {
			fmt.Printf("Continuing at signal: %s (index: %d)\n", signal, i)
			continue // Skip the rest of the loop iteration if signal is green
		}
		fmt.Printf("Stopping at signal: %s (index: %d)\n", signal, i)
		break // Exit the loop if signal is not green
	}
}
