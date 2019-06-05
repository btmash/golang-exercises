package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
	fmt.Println(&p) // See address of j (since p is set to point to j)
	fmt.Println(&i) // See address of i
	fmt.Println(&j) // See address of j
}
