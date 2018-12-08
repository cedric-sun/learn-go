package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// mapping string to Vertex
var m map[string]Vertex

func main() {
	if m == nil {
		fmt.Println("m is nil initially")
	}
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// map literal
	var anotherMap = map[string]Vertex {
		"University of Florida": Vertex{
			123,456,	// the last comma is necessary
		},
		"Google": Vertex{
			234,345,	// the last comma is necessary
		},
	}

	fmt.Println(anotherMap["University of Florida"])

	// map literal with omitted type
	var yetAnotherMap = map[string]Vertex {
		"University of Donghua": {131,797},
		"University of Maryland": {465,159},	// the last comma is necessary
	}
	fmt.Println(yetAnotherMap["University of Maryland"])

	capital := make(map[string]string)
	// put
	capital["China"] = "beijing"
	capital["USA"] = "washington"
	capital["UK"] = "london"

	// delete
	delete(capital, "UK")

	// containsKey
	elem, ok := capital["USA"]
	if ok == true {
		fmt.Println("USA is in the record")
	}

	elem, ok = capital["UK"]
	if ok == false {
		fmt.Printf("UK is not in the record, elem is: %q\n", elem)
	}
}
