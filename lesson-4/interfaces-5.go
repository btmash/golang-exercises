package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	// This fails because the M() is not implemented for i.
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
