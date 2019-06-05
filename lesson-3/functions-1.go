package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	// Compute function will run which will take a function and run it with its own set of floats inside.
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
