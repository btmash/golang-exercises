package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// Only slices can have empty 'values' and they go from position in array up to (but not including) second index.
	var s []int = primes[1:4]
	fmt.Println(s)
}
