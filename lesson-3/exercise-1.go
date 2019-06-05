package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	myArray := make([][]uint8, dy)
	for i := range myArray {
		myArray[i] = make([]uint8, dx)
		for j := range myArray[i] {
			myArray[i][j] = uint8(i * j)
		}
	}
	return myArray
}

func main() {
	pic.Show(Pic)
}
