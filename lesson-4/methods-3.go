package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// This will not modify the value of v
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Since v is a pointer, its values can be manipulated
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v)
	fmt.Println(v.Abs())
}
