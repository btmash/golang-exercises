package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}
func main() {
	firstNum := 42
	secondNum := 13
	fmt.Println(firstNum, "plus", secondNum, "equals", add(firstNum, secondNum))
}
