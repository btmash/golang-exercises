package main

import (
	"fmt"
)

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := -1
	current := 0
	sum := 0
	return func() int {
		if prev < 0 {
			prev = 0
		} else if current <= 0 {
			current = 1
			sum = 1
		} else {
			sum = current + prev
			prev = current
			current = sum
		}
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
