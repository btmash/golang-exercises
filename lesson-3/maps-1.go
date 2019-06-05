package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	m["Ashok"] = Vertex{98.6, 10.0}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m)


	var n = map[string]Vertex{
		"Bell Labs": Vertex{
			40.68433, -74.39967,
		},
		"Google": Vertex{
			37.42202, -122.08408,
		},
	}
	// An element in a hashmap cannot be modified. It can, however, be deleted or replaced.
	g, _ := n["Bell Labs"]
	g.Lat = 32
	n["Bell Labs"] = g

	fmt.Println(n)
}
