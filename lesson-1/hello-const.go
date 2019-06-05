package main

import (
	"fmt"
)

// Constants cannot be declared using := syntax
const Pi = 3.14

func main() {
	const World = "World"
	fmt.Println("Happy", Pi, "day")
	fmt.Println("Hello", World)
}
