package main

import "fmt"

type Vertex struct {
	X int
	Y int
	z bool
}

func main() {
	fmt.Println(Vertex{1, 2, false})
}
