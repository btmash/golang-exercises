package main

import (
	"fmt"
	"math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// functional programming: array_walk equivalent.
	// if you don't want to use one of k,v, use _ to kinda ignore it.
	for i, v := range pow {
		power := math.Pow(2, float64(i))
		fmt.Printf("2**%d = %d\n", i, int(power))
		if v != int(power) {
			fmt.Println("%d is not equal to %d", v, int(power))
		}
	}
}
