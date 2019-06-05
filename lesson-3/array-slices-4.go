package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	var m = s[1:4]
	fmt.Println(m)

	var n = s[:2]
	fmt.Println(n)

	var o = s[1:]
	fmt.Println(o)

	var p []int = s[:]
	fmt.Println(p)
}
