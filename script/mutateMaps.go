package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

//	<sting, Vetex>
// If the top-level type is just a type name, you can omit it from the elements of the literal.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}


func main() {
	m := make(map[string]int)
	
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	//containsKey
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
