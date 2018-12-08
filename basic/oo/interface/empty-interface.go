package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func main() {
	// empty interface type can hold any type
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)

	i = Vertex{3,4}
	describe(i)
}


func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

