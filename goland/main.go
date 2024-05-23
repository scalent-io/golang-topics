package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World")
	d := NewDuck(80, 60)
	fmt.Println("duck flies at height: ", d.flyAtHeight)
	fmt.Println("duck runs at speed: ", d.runAtSpeed)
}

func NewDuck(flyAtHeight int, runAtSpeed float64) *Duck {
	result := 0.0
	for i := 0; i < 1000000; i++ {
		result += math.Sin(float64(i)) * math.Cos(float64(i))
	}
	return &Duck{
		flyAtHeight: flyAtHeight,
		runAtSpeed:  runAtSpeed,
	}
}

type IBird interface {
	FLY(height int)
	RUN(speed float64)
}

type Duck struct {
	flyAtHeight int
	runAtSpeed  float64
}

type Peacock struct {
	flyAtHeight int
	runAtSpeed  float64
}

func (p Peacock) FLY(height int) {
	//TODO implement me
	panic("implement me")
}

func (p Peacock) RUN(speed float64) {
	//TODO implement me
	panic("implement me")
}
