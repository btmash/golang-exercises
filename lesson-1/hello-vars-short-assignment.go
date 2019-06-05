package main

import "fmt"

func main() {
	var i, j int = 3, 6
	// Short assignment is only available inside functions. Not outside them.
	k:= 3
	c, python, java := "see", false, true
	fmt.Println(i, j, k, c, python, java)
}
