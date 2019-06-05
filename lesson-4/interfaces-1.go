package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	// FIRST SET VARIABLE TO INTERFACE. THEN SET VARIABLE TO TYPE THAT IMPLEMENTS SAID INTERFACE
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// We can implement it not as *Vertex while would let us use either format
	// Though we lose on the capability to modify Vertex.
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

//func (v *Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}

func (v Vertex) Abs() float64 {
	return math.Abs(v.X + v.Y)
}
