package main

import (
	"fmt"
	"math"
)

func powa(x, n, lim float64) float64 {
	// Declarations allowed will still work.
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		powa(3, 2, 10),
		powa(3, 3, 20),
	)
}
