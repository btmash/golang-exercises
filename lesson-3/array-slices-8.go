package main

import "fmt"

func main() {
	// zero-ed array of length 5
	a := make([]int, 5)
	printSlice("a", a)

	// unfilled array of length 5 but zero-ed array is allocated
	b := make([]int, 0, 5)
	printSlice("b", b)

	// len 2 cap 2
	c := b[:2]
	printSlice("c", c)

	// len 0 cap 3
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
