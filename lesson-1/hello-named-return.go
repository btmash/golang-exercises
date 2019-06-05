package main

import (
	"fmt"
)

// Preferably, do not use this approach. Be explicit in what you are returning on the return statement otherwise
// readability is going to be shit, quite frankly.
func summation(x, y int) (sum int) {
	sum = x + y
	return
}
func main() {
	firstNum := 77
	secondNum := 14
	fmt.Println(firstNum, "plus", secondNum, "equals", summation(firstNum, secondNum))
}
