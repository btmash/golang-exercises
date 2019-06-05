package main

import "fmt"

func check() {
	fmt.Println("ashok")
}

func main() {
	defer check()
	defer fmt.Println("world")
	fmt.Println("hello")
}
