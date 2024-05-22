package main

import (
	"fmt"
)

type IntStruct struct {
	Field1 int
	Field2 int
}

type FloatStruct struct {
	Field1 float64
	Field2 float64
}

func main() {
	intStruct := IntStruct{Field1: 1, Field2: 2}
	floatStruct := FloatStruct{
		Field1: float64(intStruct.Field1),
		Field2: float64(intStruct.Field2),
	}

	fmt.Println("Converted float struct:", floatStruct)
}
