package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 1; i <= 10; i++ {
		z = z - ((z*z - x) / (2*z))
		fmt.Println(i, z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(81))
	fmt.Println(math.Sqrt(81))
}
