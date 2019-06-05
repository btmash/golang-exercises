package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Note the (v Vertex) before function name. Think if it 'attaching' itself to the Vertex struct.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Same as non struct version below which now has to take arguments.
func Nabs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Nabs(v))
}
