package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 10
	fmt.Println(v.X, v.Y)
}
