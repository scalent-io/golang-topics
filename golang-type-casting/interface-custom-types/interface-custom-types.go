package main

import (
	"fmt"
)

type Stringer interface {
	String() string
}

type MyType struct {
	value string
}

func (m MyType) String() string {
	return m.value
}

func main() {
	var s Stringer = MyType{"Hello, Go!"}

	// Type assertion to convert Stringer to MyType
	if myType, ok := s.(MyType); ok {
		fmt.Println("The MyType value is:", myType.value)
	} else {
		fmt.Println("Type assertion failed")
	}
}
