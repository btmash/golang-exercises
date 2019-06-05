package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// This will reference names[1], a[1], b[0] since its a reference to the original.
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
