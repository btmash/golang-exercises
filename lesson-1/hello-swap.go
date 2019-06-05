package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	firstString := "hello"
	secondString := "world"
	fmt.Println(swap(firstString, secondString))
}
