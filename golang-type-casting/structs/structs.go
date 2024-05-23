package main

import (
	"fmt"
)

type IntStruct struct {
	Number1 int
	Number2 int
}

type FloatStruct struct {
	Number1 float64
	Number2 float64
}

func main() {
	intStruct := IntStruct{Number1: 121, Number2: 144}
	floatStruct := FloatStruct{
		Number1: float64(intStruct.Number1),
		Number2: float64(intStruct.Number2),
	}

	fmt.Println("Converted float struct:", floatStruct)
}
